package yml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
	"os"
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
	sourceYml := map[interface{}]interface{}{}
	for _, file := range p.files {
		temp, err := p.load(file)
		if err != nil {
			return err
		}
		if err := p.appendUpdate(sourceYml, temp); err != nil {
			return err
		}
	}
	out, err := yaml.Marshal(&sourceYml)
	if err != nil {
		return err
	} else {
		ioutil.WriteFile(p.target, out, os.ModePerm)
	}
	return nil
}

func (p *Merger) load(file string) (map[interface{}]interface{}, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := map[interface{}]interface{}{}
	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (p *Merger) appendUpdate(target map[interface{}]interface{}, source map[interface{}]interface{}) error {
	for k, v := range source {
		tv := target[k]
		if tv != nil {
			sourceMap, sok := v.(map[interface{}]interface{})
			if sok {
				targetMap, tok := tv.(map[interface{}]interface{})
				if tok {
					p.appendUpdate(targetMap, sourceMap)
				} else {
					return fmt.Errorf("target %s isn't map", k)
				}
			} else {
				sourceArray, sok := v.([]interface{})
				if sok {
					targetArray, tok := tv.([]interface{})
					if tok {
						for _, item := range sourceArray {
							targetArray = append(targetArray, item)
						}
						target[k] = targetArray
					} else {
						return fmt.Errorf("target %s isn't array", k)
					}
				} else {
					target[k] = v
				}
			}

		} else {
			target[k] = v
		}
	}

	return nil
}

//func (p *Merger) appendUpdateDelete(target map[interface{}]interface{}, source map[interface{}]interface{}) error {
//	removed := []interface{}{}
//	for k, v := range target {
//		sv := source[k]
//		if sv != nil {
//			targetMap, tok := v.(map[interface{}]interface{})
//			if tok {
//				sourceMap, sok := sv.(map[interface{}]interface{})
//				if sok {
//					p.appendUpdateDelete(targetMap, sourceMap)
//				} else {
//					return fmt.Errorf("the type of %s in source isn't map", k)
//				}
//			} else {
//				targetArray, tok := v.([]interface{})
//				if tok {
//					sourceArray, sok := sv.([]interface{})
//					if sok {
//
//					} else {
//						return fmt.Errorf("the type of %s in source isn't array", k)
//					}
//				} else {
//					target[k] = sv
//				}
//			}
//
//		} else {
//			removed = append(removed, k)
//		}
//	}
//
//	for _, k := range removed {
//		delete(target, k)
//	}
//	return nil
//}

