package util

import (
	"fmt"
	"reflect"
)

func SetValue(o interface{}, attrs []string, values []interface{}) error {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("o isn't Ptr")
		if !v.Elem().CanSet() {
			return fmt.Errorf("o can't set")
		}
	}
	v = v.Elem()
	for i, attr := range attrs {
		fv := values[i]
		f := v.FieldByName(attr)
		if f.IsValid() {
			switch f.Kind() {
			case reflect.String:
				setString(f, fv)
			case reflect.Bool:
				setBool(f, fv)
			case reflect.Uint:
				setUint(f, fv)
			case reflect.Uint8:
				setUint(f, fv)
			case reflect.Uint16:
				setUint(f, fv)
			case reflect.Uint32:
				setUint(f, fv)
			case reflect.Uint64:
				setUint(f, fv)
			case reflect.Int:
				setInt(f, fv)
			case reflect.Int8:
				setInt(f, fv)
			case reflect.Int16:
				setInt(f, fv)
			case reflect.Int32:
				setInt(f, fv)
			case reflect.Int64:
				setInt(f, fv)
			case reflect.Float32:
				setFloat(f, fv)
			case reflect.Float64:
				setFloat(f, fv)
			}
		}
	}
	return nil
}

func setString(f reflect.Value, v interface{}) {
	if s, ok := v.(string); ok {
		f.SetString(s)
	}
}

func setBool(f reflect.Value, v interface{}) {
	if s, ok := v.(bool); ok {
		f.SetBool(s)
	}
}

func setUint(f reflect.Value, v interface{}) {
	if s, ok := v.(uint64); ok {
		f.SetUint(s)
	}
}

func setInt(f reflect.Value, v interface{}) {
	if s, ok := v.(int64); ok {
		f.SetInt(s)
	}
}

func setFloat(f reflect.Value, v interface{}) {
	if s, ok := v.(float64); ok {
		f.SetFloat(s)
	}
}
