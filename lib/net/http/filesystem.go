// HTTP file system request handler
package http

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"net/http"
	"github.com/pressly/chi/render"
	"encoding/json"
//	"log"
)

const sniffLen = 512

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

type PutRequest struct {
	Cmd  string `json:"cmd"`
	Data *json.RawMessage    `json:"data"`
}

type File struct {
	Url  string        `json:"url"`
	Name string        `json:"name"`
	Type string        `json:"type"`
}

type FileSystem interface {
	Mkdir(name string) (http.File, error)
	Open(name string) (http.File, error)
	Create(name string) (http.File, error)
	Delete(name string) error
	Rename(old string, new string) error
	Update(name string, content string) error
}

// A Dir implements FileSystem using the native file system restricted to a
// specific directory tree.
//
// While the FileSystem.Open method takes '/'-separated paths, a Dir's string
// value is a filename on the native file system, not a URL, so it is separated
// by filepath.Separator, which isn't necessarily '/'.
//
// An empty Dir is treated as ".".
type Dir string

func (d Dir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
			strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	f, err := os.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name))))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (d Dir) Mkdir(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
			strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	path := filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name)))
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return d.Open(name)
}

func (d Dir) Create(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
			strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	f, err := os.Create(filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name))))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (d Dir) Delete(name string) error {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
			strings.Contains(name, "\x00") {
		return errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	err := os.RemoveAll(filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name))))
	if err != nil {
		return err
	}
	return nil
}

func (d Dir) Rename(name string, new string) error {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
			strings.Contains(name, "\x00") {
		return errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	oldPath := filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name)))
	newPath := filepath.Join(dir, filepath.FromSlash(path.Clean("/" + new)))
	//log.Println(oldPath)
	//log.Println(newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}
	return nil
}

func (d Dir) Update(name string, content string) error {
	//if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
	//		strings.Contains(name, "\x00") {
	//	return errors.New("http: invalid character in file path")
	//}
	//dir := string(d)
	//if dir == "" {
	//	dir = "."
	//}
	//oldPath := filepath.Join(dir, filepath.FromSlash(path.Clean("/" + name)))
	//newPath := filepath.Join(dir, filepath.FromSlash(path.Clean("/" + new)))
	//err := os.(oldPath, newPath)
	//if err != nil {
	//	return err
	//}
	return nil
}

//// A FileSystem implements access to a collection of named files.
//// The elements in a file path are separated by slash ('/', U+002F)
//// characters, regardless of host operating system convention.
//type FileSystem interface {
//	Open(name string) (File, error)
//}

// A File is returned by a FileSystem's Open method and can be
// served by the FileServer implementation.
//
// The methods should behave the same as those on an *os.File.
//type File interface {
//	io.Closer
//	io.Reader
//	io.Seeker
//	Readdir(count int) ([]os.FileInfo, error)
//	Stat() (os.FileInfo, error)
//}

func dirList(f http.File, path string) ([]File, error) {
	dirs, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	sort.Sort(byName(dirs))
	files := []File{}
	for _, d := range dirs {
		name := d.Name()
		if name[0] == '.' {
			continue
		}
		ftype := "file"
		if d.IsDir() {
			ftype = "dir"
		}

		url := url.URL{Path: path + name}
		file := File{Url:url.String(), Name:name, Type: ftype}
		files = append(files, file)
	}
	return files, nil
}


// if name is empty, filename is unknown. (used for mime type, before sniffing)
// if modtime.IsZero(), modtime is unknown.
// content must be seeked to the beginning of the file.
// The sizeFunc is called at most once. Its error, if any, is sent in the HTTP response.
func serveContent(w http.ResponseWriter, r *http.Request, name string, modtime time.Time, sizeFunc func() (int64, error), content io.ReadSeeker) {
	if checkLastModified(w, r, modtime) {
		return
	}
	rangeReq, done := checkETag(w, r, modtime)
	if done {
		return
	}

	code := http.StatusOK

	// If Content-Type isn't set, use the file's extension to find it, but
	// if the Content-Type is unset explicitly, do not sniff the type.
	ctypes, haveType := w.Header()["Content-Type"]
	var ctype string
	if !haveType {
		ctype = mime.TypeByExtension(filepath.Ext(name))
		if ctype == "" {
			// read a chunk to decide between utf-8 text and binary
			var buf [sniffLen]byte
			n, _ := io.ReadFull(content, buf[:])
			ctype = http.DetectContentType(buf[:n])
			_, err := content.Seek(0, io.SeekStart) // rewind to output whole file
			if err != nil {
				http.Error(w, "seeker can't seek", http.StatusInternalServerError)
				return
			}
		}
		w.Header().Set("Content-Type", ctype)
	} else if len(ctypes) > 0 {
		ctype = ctypes[0]
	}

	size, err := sizeFunc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// handle Content-Range header.
	sendSize := size
	var sendContent io.Reader = content
	if size >= 0 {
		ranges, err := parseRange(rangeReq, size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusRequestedRangeNotSatisfiable)
			return
		}
		if sumRangesSize(ranges) > size {
			// The total number of bytes in all the ranges
			// is larger than the size of the file by
			// itself, so this is probably an attack, or a
			// dumb client. Ignore the range request.
			ranges = nil
		}
		switch {
		case len(ranges) == 1:
			// RFC 2616, Section 14.16:
			// "When an HTTP message includes the content of a single
			// range (for example, a response to a request for a
			// single range, or to a request for a set of ranges
			// that overlap without any holes), this content is
			// transmitted with a Content-Range header, and a
			// Content-Length header showing the number of bytes
			// actually transferred.
			// ...
			// A response to a request for a single range MUST NOT
			// be sent using the multipart/byteranges media type."
			ra := ranges[0]
			if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
				http.Error(w, err.Error(), http.StatusRequestedRangeNotSatisfiable)
				return
			}
			sendSize = ra.length
			code = http.StatusPartialContent
			w.Header().Set("Content-Range", ra.contentRange(size))
		case len(ranges) > 1:
			sendSize = rangesMIMESize(ranges, ctype, size)
			code = http.StatusPartialContent

			pr, pw := io.Pipe()
			mw := multipart.NewWriter(pw)
			w.Header().Set("Content-Type", "multipart/byteranges; boundary=" + mw.Boundary())
			sendContent = pr
			defer pr.Close() // cause writing goroutine to fail and exit if CopyN doesn't finish.
			go func() {
				for _, ra := range ranges {
					part, err := mw.CreatePart(ra.mimeHeader(ctype, size))
					if err != nil {
						pw.CloseWithError(err)
						return
					}
					if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
						pw.CloseWithError(err)
						return
					}
					if _, err := io.CopyN(part, content, ra.length); err != nil {
						pw.CloseWithError(err)
						return
					}
				}
				mw.Close()
				pw.Close()
			}()
		}

		w.Header().Set("Accept-Ranges", "bytes")
		if w.Header().Get("Content-Encoding") == "" {
			w.Header().Set("Content-Length", strconv.FormatInt(sendSize, 10))
		}
	}

	w.WriteHeader(code)

	if r.Method != "HEAD" {
		io.CopyN(w, sendContent, sendSize)
	}
}

var unixEpochTime = time.Unix(0, 0)

// modtime is the modification time of the resource to be served, or IsZero().
// return value is whether this request is now complete.
func checkLastModified(w http.ResponseWriter, r *http.Request, modtime time.Time) bool {
	if modtime.IsZero() || modtime.Equal(unixEpochTime) {
		// If the file doesn't have a modtime (IsZero), or the modtime
		// is obviously garbage (Unix time == 0), then ignore modtimes
		// and don't process the If-Modified-Since header.
		return false
	}

	// The Date-Modified header truncates sub-second precision, so
	// use mtime < t+1s instead of mtime <= t to check for unmodified.
	if t, err := time.Parse(http.TimeFormat, r.Header.Get("If-Modified-Since")); err == nil && modtime.Before(t.Add(1 * time.Second)) {
		h := w.Header()
		delete(h, "Content-Type")
		delete(h, "Content-Length")
		w.WriteHeader(http.StatusNotModified)
		return true
	}
	w.Header().Set("Last-Modified", modtime.UTC().Format(http.TimeFormat))
	return false
}

// checkETag implements If-None-Match and If-Range checks.
//
// The ETag or modtime must have been previously set in the
// ResponseWriter's headers. The modtime is only compared at second
// granularity and may be the zero value to mean unknown.
//
// The return value is the effective request "Range" header to use and
// whether this request is now considered done.
func checkETag(w http.ResponseWriter, r *http.Request, modtime time.Time) (rangeReq string, done bool) {
	etag := getHeader(w.Header(), "Etag")
	rangeReq = getHeader(r.Header, "Range")

	// Invalidate the range request if the entity doesn't match the one
	// the client was expecting.
	// "If-Range: version" means "ignore the Range: header unless version matches the
	// current file."
	// We only support ETag versions.
	// The caller must have set the ETag on the response already.
	if ir := getHeader(r.Header, "If-Range"); ir != "" && ir != etag {
		// The If-Range value is typically the ETag value, but it may also be
		// the modtime date. See golang.org/issue/8367.
		timeMatches := false
		if !modtime.IsZero() {
			if t, err := http.ParseTime(ir); err == nil && t.Unix() == modtime.Unix() {
				timeMatches = true
			}
		}
		if !timeMatches {
			rangeReq = ""
		}
	}

	if inm := getHeader(r.Header, "If-None-Match"); inm != "" {
		// Must know ETag.
		if etag == "" {
			return rangeReq, false
		}

		// TODO(bradfitz): non-GET/HEAD requests require more work:
		// sending a different status code on matches, and
		// also can't use weak cache validators (those with a "W/
		// prefix).  But most users of ServeContent will be using
		// it on GET or HEAD, so only support those for now.
		if r.Method != "GET" && r.Method != "HEAD" {
			return rangeReq, false
		}

		// TODO(bradfitz): deal with comma-separated or multiple-valued
		// list of If-None-match values. For now just handle the common
		// case of a single item.
		if inm == etag || inm == "*" {
			h := w.Header()
			delete(h, "Content-Type")
			delete(h, "Content-Length")
			w.WriteHeader(http.StatusNotModified)
			return "", true
		}
	}
	return rangeReq, false
}

// name is '/'-separated, not filepath.Separator.
func serveFile(w http.ResponseWriter, r *http.Request, fs FileSystem, name string) {
	f, err := fs.Open(name)
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}

	// Still a directory? (we didn't find an index.html file)
	if d.IsDir() {
		path := name
		if path != "/" {
			path = path + "/"
		}
		files, err := dirList(f, path)
		if err != nil {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error reading directory")
		} else {
			buf, err := json.Marshal(files)
			if err != nil {
				w.Header().Set("Content-Type", "text/json; charset=utf-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error json marshal")
			} else {
				w.Header().Set("Content-Type", "text/json; charset=utf-8")
				w.Write(buf)
			}

		}
		return
	}

	// serveContent will check modification time
	sizeFunc := func() (int64, error) {
		return d.Size(), nil
	}
	serveContent(w, r, d.Name(), d.ModTime(), sizeFunc, f)
}

func createFile(w http.ResponseWriter, r *http.Request, fs FileSystem, name string) {
	index := strings.LastIndex(name, "/")
	fname := name[index + 1:]
	path := name
	if path != "/" {
		path = path + "/"
	}

	if strings.LastIndex(name, ".") > 0 {
		f, err := fs.Create(name)
		if err != nil {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}
		defer f.Close()

		file := File{Url:name, Name:fname, Type: "file"}

		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		buf, err := json.Marshal(file)
		if err != nil {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error json marshal")
		} else {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Write(buf)
		}

	} else {
		f, err := fs.Mkdir(name)
		if err != nil {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}
		defer f.Close()

		file := File{Url:name, Name:fname, Type: "dir"}

		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		buf, err := json.Marshal(file)
		if err != nil {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error json marshal")
		} else {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Write(buf)
		}
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request, fs FileSystem, name string) {
	if err := fs.Delete(name); err != nil {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		fmt.Fprintf(w, "{\"url\":\"%s\"}", name)
	}
}

func renameFile(w http.ResponseWriter, r *http.Request, fs FileSystem, old string, new string) {
	if err := fs.Rename(old, new); err != nil {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	} else {
		w.Header().Set("Content-Type", "text/json; charset=utf-8")
		fmt.Fprintf(w, "{\"url\":\"%s\"}", new)
	}
}

func updateFile(w http.ResponseWriter, r *http.Request, fs FileSystem, name string, file *File) {
	if name != file.Name {
		//update name
		if err := fs.Rename(name, file.Url); err != nil {
			render.Status(r, 500)
			render.JSON(w, r, err.Error())
		} else {
			render.Status(r, 200)
			render.JSON(w, r, file)
		}
	} else {
		//update file
	}
	//if err := fs.Update(name, content); err != nil {
	//	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintln(w, err.Error())
	//} else {
	//	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	//	fmt.Fprintf(w, "{\"url\":\"%s\"}", new)
	//}
}

// toHTTPError returns a non-specific HTTP error message and status code
// for a given non-nil error value. It's important that toHTTPError does not
// actually return err.Error(), since msg and httpStatus are returned to users,
// and historically Go's ServeContent always returned just "404 Not Found" for
// all errors. We don't want to start leaking information in error messages.
func toHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404 page not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}

type FileHandler struct {
	root FileSystem
}

// FileServer returns a handler that serves HTTP requests
// with the contents of the file system rooted at root.
//
// To use the operating system's file system implementation,
// use http.Dir:
//
//     http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// As a special case, the returned file server redirects any request
// ending in "/index.html" to the same path, without the final
// "index.html".
func FileServer(root FileSystem) *FileHandler {
	return &FileHandler{root}
}

func (f *FileHandler) Get(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	serveFile(w, r, f.root, path.Clean(upath))
}

func (f *FileHandler) Post(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	createFile(w, r, f.root, path.Clean(upath))
}

func (f *FileHandler) Put(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	file := &File{}
	if err := render.Bind(r.Body, file); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, err.Error())
		return
	}
	updateFile(w, r, f.root, path.Clean(upath), file)
}

func (f *FileHandler) Delete(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	deleteFile(w, r, f.root, path.Clean(upath))
}

// httpRange specifies the byte range to be sent to the client.
type httpRange struct {
	start, length int64
}

func (r httpRange) contentRange(size int64) string {
	return fmt.Sprintf("bytes %d-%d/%d", r.start, r.start + r.length - 1, size)
}

func (r httpRange) mimeHeader(contentType string, size int64) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		"Content-Range": {r.contentRange(size)},
		"Content-Type":  {contentType},
	}
}

// parseRange parses a Range header string as per RFC 2616.
func parseRange(s string, size int64) ([]httpRange, error) {
	if s == "" {
		return nil, nil // header not present
	}
	const b = "bytes="
	if !strings.HasPrefix(s, b) {
		return nil, errors.New("invalid range")
	}
	var ranges []httpRange
	for _, ra := range strings.Split(s[len(b):], ",") {
		ra = strings.TrimSpace(ra)
		if ra == "" {
			continue
		}
		i := strings.Index(ra, "-")
		if i < 0 {
			return nil, errors.New("invalid range")
		}
		start, end := strings.TrimSpace(ra[:i]), strings.TrimSpace(ra[i + 1:])
		var r httpRange
		if start == "" {
			// If no start is specified, end specifies the
			// range start relative to the end of the file.
			i, err := strconv.ParseInt(end, 10, 64)
			if err != nil {
				return nil, errors.New("invalid range")
			}
			if i > size {
				i = size
			}
			r.start = size - i
			r.length = size - r.start
		} else {
			i, err := strconv.ParseInt(start, 10, 64)
			if err != nil || i >= size || i < 0 {
				return nil, errors.New("invalid range")
			}
			r.start = i
			if end == "" {
				// If no end is specified, range extends to end of the file.
				r.length = size - r.start
			} else {
				i, err := strconv.ParseInt(end, 10, 64)
				if err != nil || r.start > i {
					return nil, errors.New("invalid range")
				}
				if i >= size {
					i = size - 1
				}
				r.length = i - r.start + 1
			}
		}
		ranges = append(ranges, r)
	}
	return ranges, nil
}

// countingWriter counts how many bytes have been written to it.
type countingWriter int64

func (w *countingWriter) Write(p []byte) (n int, err error) {
	*w += countingWriter(len(p))
	return len(p), nil
}

// rangesMIMESize returns the number of bytes it takes to encode the
// provided ranges as a multipart response.
func rangesMIMESize(ranges []httpRange, contentType string, contentSize int64) (encSize int64) {
	var w countingWriter
	mw := multipart.NewWriter(&w)
	for _, ra := range ranges {
		mw.CreatePart(ra.mimeHeader(contentType, contentSize))
		encSize += ra.length
	}
	mw.Close()
	encSize += int64(w)
	return
}

func sumRangesSize(ranges []httpRange) (size int64) {
	for _, ra := range ranges {
		size += ra.length
	}
	return
}

type byName []os.FileInfo

func (s byName) Len() int {
	return len(s)
}
func (s byName) Less(i, j int) bool {
	return s[i].Name() < s[j].Name()
}
func (s byName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getHeader(h http.Header, key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}
