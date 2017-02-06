package file

import (
	"net/http"
	"os"
	"io"
	"github.com/pressly/chi/render"

	"github.com/pressly/chi"
	"strings"
	"path/filepath"
	"path"
)

func UploadServer(mx *chi.Mux, path string, fileRoot string) {
	if strings.ContainsAny(path, ":*") {
		panic("chi: FileServer does not permit URL parameters.")
	}
	upload := Upload(fileRoot)
	prefix := path
	path += "*"

	mx.Post(path, exec(prefix, func(w http.ResponseWriter, r *http.Request) {
		upload.Post(w, r)
	}))
}

func exec(prefix string, fn http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r.URL.Path = p
			fn(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}

type Upload string

func (p Upload) Post(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("upload: %s %s\n", req.URL.Path, p)
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("uploadfile")
	if err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err)
		return
	}
	defer file.Close()
	//fmt.Println(handler.Filename)
	path := filepath.Join(string(p), filepath.FromSlash(path.Clean("/" + req.URL.Path)))
	err = os.MkdirAll(path, 0777)
	if err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err)
		return
	}

	f, err := os.OpenFile(path + "/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
