package util

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ChaosXu/nerv/lib/env"
	"github.com/toolkits/file"
	"log"
)

func LoadScript(scriptUrl string) (string, error) {
	dir := env.Config().GetMapString("scripts", "cache")
	if dir == "" {
		return "", fmt.Errorf("scripts.cache isn't setted");
	}
	url, err := url.Parse(scriptUrl)
	if err != nil {
		return "", fmt.Errorf("error url %s", url);
	}
	l := strings.LastIndex(url.Path, "/")
	dir = dir + url.Path[:l]
	script := dir + url.Path[l:]
	log.Println(scriptUrl)
	scriptContent, err := file.ToString(script)
	if err != nil {
		os.MkdirAll(dir, os.ModeDir | os.ModePerm)
		if err := file.Download(script, scriptUrl); err != nil {
			return "", err
		}
		return file.ToString(script)
	} else {
		return scriptContent, nil
	}

}
