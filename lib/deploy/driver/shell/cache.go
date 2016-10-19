package shell

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ChaosXu/nerv/lib/env"
	"github.com/toolkits/file"
)

func loadScript(scriptUrl string) (string, error) {
	dir := env.Config().GetMapString("cache", "scripts")
	if dir == "" {
		return "", fmt.Errorf("cache.script isn't setted");
	}
	url, err := url.Parse(scriptUrl)
	if err != nil {
		return "", fmt.Errorf("error url %s", url);
	}
	l := strings.LastIndex(url.Path, "/")
	dir = dir + url.Path[:l]
	script := dir + url.Path[l:]
	os.MkdirAll(dir, os.ModeDir | os.ModePerm)
	if err := file.Download(script, scriptUrl); err != nil {
		return "", err
	}

	return file.ToString(script)
}
