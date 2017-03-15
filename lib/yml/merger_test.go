package yml

import (
	"testing"
	"io/ioutil"
	"fmt"
	"reflect"
	"gopkg.in/yaml.v2"
)

func TestMerger(t *testing.T) {

	merger := NewMerger("merged.yml")
	merger.Add("source1.yml")
	merger.Add("source2.yml")
	merger.Merge()

	t2, err := loadFile("merged.yml")
	if err != nil {
		t.Error(err)
	}

	t1, err := loadFile("target.yml")
	if err != nil {
		t.Error(err)
	}

	//if !reflect.DeepEqual(t1,t2) {
	//
	//}
	if err := equals(t1, t2); err != nil {
		t.Error(err)
	}
}

func loadFile(file string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	out := map[string]interface{}{}
	if err := yaml.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func equals(t map[string]interface{}, s map[string]interface{}) error {
	for tk, tv := range t {
		sv := s[tk]
		if sv == nil {
			return fmt.Errorf("could not found %s in source map", tk)
		}

		tmap, tok := tv.(map[string]interface{})
		if tok {
			smap, sok := sv.(map[string]interface{})
			if sok {
				if err := equals(tmap, smap); err != nil {
					return fmt.Errorf("%s. parent:%s", err.Error(), tk)
				}
			} else {
				return fmt.Errorf("the type of %s isn't equals", tk)
			}
		} else {
			tarray, tok := tv.([]interface{})
			if tok {
				sarray, sok := sv.([]interface{})
				if sok {
					if err := equalsArray(tarray, sarray); err != nil {
						return fmt.Errorf("%s. parent:%s", err.Error(), tk)
					}
				}
			} else {
				if !equalsValue(tv, sv) {
					return fmt.Errorf("don't equals. index:%v", tk)
				}
			}
		}
	}
	return nil
}

func equalsArray(t []interface{}, s []interface{}) error {
	if len(t) != len(s) {
		return fmt.Errorf("target len is %v,source len is %v", len(t), len(s))
	}

	for i, tv := range t {
		sv := s[i]
		tmap, tok := tv.(map[string]interface{})
		if tok {
			smap, sok := sv.(map[string]interface{})
			if sok {
				if err := equals(tmap, smap); err != nil {
					return fmt.Errorf("%s,index:%v", err.Error(), i)
				}
			} else {
				return fmt.Errorf("source isn't map,index:%v", i)
			}
		} else {
			tarray, tok := tv.([]interface{})
			if tok {
				sarray, sok := sv.([]interface{})
				if sok {
					if err := equalsArray(tarray, sarray); err != nil {
						return fmt.Errorf("%s. index:%v", err.Error(), i)
					}
				}
			} else {
				if !equalsValue(tv, sv) {
					return fmt.Errorf("don't equals. index:%v", i)
				}
			}
		}
	}
	return nil
}

func equalsValue(t, s interface{}) bool {
	return reflect.DeepEqual(t, s)
}
