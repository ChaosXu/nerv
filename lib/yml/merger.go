package yml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"fmt"
)

// Merger merge multi yml file in one
type Merger struct {
	target string
	files  []string
}

func NewMerger(target string) *Merger {
	return &Merger{target:target, files:[]string{}}
}

// Add a file to merge
func (p *Merger) Add(file string) {
	p.files = append(p.files, file)
}

// Remove a file then remove the file's elements from target file
// return a removed element or "" if not to be removed
func (p *Merger) Remove(file string) string {
	for i, f := range p.files {
		if f == file {
			removed := p.files[i]
			p.files = append(p.files[:i], p.files[i + 1:]...)
			return removed
		}
	}
	return ""
}

// Merge all files into one
func (p *Merger) Merge() error {
	sourceYml := map[string]interface{}{}
	for _, file := range p.files {
		temp, err := p.load(file)
		if err != nil {
			return err
		}
		if err := p.appendUpdate(sourceYml, temp); err != nil {
			return err
		}
	}
	targetYml, err := p.load(p.target)
	if err != nil {
		return err
	}
	p.appendUpdateDelete(targetYml, sourceYml)
	return nil
}

func (p *Merger) load(file string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := &map[string]interface{}{}
	if err := yaml.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (p *Merger) appendUpdate(target map[string]interface{}, source map[string]interface{}) error {
	for k, v := range target {
		sv := source[k]
		if sv != nil {
			targetMap, ok := v.(map[string]interface{})
			if !ok {
				target[k] = sv
			}
			sourceMap, ok := sv.(map[string]interface{})
			if !ok {
				return fmt.Errorf("the type of %s in source isn't map", k)
			}
			p.appendUpdate(targetMap, sourceMap)
		}
	}
	return nil
}

func (p *Merger) appendUpdateDelete(target map[string]interface{}, source map[string]interface{}) error {
	removed := []string{}
	for k, v := range target {
		sv := source[k]
		if sv != nil {
			targetMap, ok := v.(map[string]interface{})
			if !ok {
				target[k] = sv
			}
			sourceMap, ok := sv.(map[string]interface{})
			if !ok {
				return fmt.Errorf("the type of %s in source isn't map", k)
			}
			p.appendUpdateDelete(targetMap, sourceMap)
		} else {
			removed = append(removed, k)
		}
	}

	for _, k := range removed {
		delete(target, k)
	}
	return nil
}

