package structCvter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// settings
var nickKey string
var hideKey string
var used = false
var firstLetterLower = false

// setter interfaces

/* set firstLetterLowerCase, defult is false*/
func SetFirstLetterLowerCase() error {
	if used == true {
		return fmt.Errorf("setting must be set before use")
	}
	firstLetterLower = true
	return nil
}

func SetHideKey(key string) error {
	if used == true {
		return fmt.Errorf("setting must be set before use")
	}
	hideKey = key
	return nil
}

/* set property nick name key*/
func SetNickKey(key string) error {
	if len(nickKey) > 0 {
		return fmt.Errorf("nick key must set only once")
	}

	if used == true {
		return fmt.Errorf("nick key must be set before use")
	}
	nickKey = key
	return nil
}

func ToJson(in interface{}) (string, error) {
	_map, err := ToMap(in)
	if err != nil {
		return "", err
	}
	data, err2 := json.Marshal(_map)
	if err2 != nil {
		return "", err2
	}
	return string(data), nil
}

func ToMap(in interface{}) (map[string]interface{}, error) {
	if used != true {
		used = true
	}
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() == reflect.Struct {
		return structToMap(in)
	} else if v.Kind() == reflect.Map {
		return mapToMap(in)
	}
	return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
}

func mapToMap(in interface{}) (map[string]interface{}, error) {
	_map := in.(map[string]interface{})
	out := make(map[string]interface{})
	for key, value := range _map {
		ty := reflect.TypeOf(value).Kind()
		var err error
		var ret map[string]interface{}
		if ty == reflect.Struct {
			ret, err = structToMap(value)
		} else if ty == reflect.Map {
			ret, err = mapToMap(value)
		}
		if err != nil {
			return nil, err
		}
		if ret != nil {
			out[key] = ret
		} else {
			out[key] = value
		}
	}
	return out, nil
}

func structToMap(in interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typ := v.Type()
	out := make(map[string]interface{})
	fieldCount := v.NumField()
	for i := 0; i < fieldCount; i++ {
		// gets us a StructField
		fi := typ.Field(i)
		vl := v.Field(i)

		if checkHidden(fi) {
			continue
		}

		if vl.CanInterface() == false {
			continue
		}
		key := getKey(fi)
		//	fmt.Println("key:", key)

		value, err := getValue(fi, vl)
		if err != nil {
			return nil, err
		}

		//	fmt.Println("value:", value)
		if value != nil {
			out[key] = value
		}

	}
	return out, nil
}

func getValue(fi reflect.StructField, v reflect.Value) (interface{}, error) {

	kind := v.Type().Kind()
	var inface interface{}
	if kind == reflect.Struct {
		inface = v.Interface()
		_map, err := ToMap(inface)
		return _map, err
	} else if kind == reflect.Map {
		inface = v.Interface()
		_map, err := ToMap(inface)
		fmt.Println("map Type convert ", _map)
		return _map, err
	} else if kind == reflect.Int {
		inface = v.Interface()
	} else if kind == reflect.Float32 {
		inface = v.Interface()
	} else if kind == reflect.String {
		inface = v.Interface()
	} else {
		inface = v.Interface()
	}

	return inface, nil
}

func checkHidden(fi reflect.StructField) bool {
	if len(hideKey) == 0 {
		return false
	}

	key := fi.Tag.Get(hideKey)
	if key == "true" {
		return true
	}

	return false
}

func getKey(fi reflect.StructField) string {
	nick := ""
	if len(nickKey) > 0 {
		nick = fi.Tag.Get(nickKey)
	}

	if len(nick) > 0 { // if has a nick name
		return nick
	}
	//  does`t has  nick name
	name := fi.Name
	// change the first letter to lower case
	if firstLetterLower {
		length := len(name)
		first := name[0:1]
		rest := ""
		restLength := length - 1
		if restLength > 0 {
			rest = name[1:length]
		}
		name = strings.ToLower(first) + rest
	}
	return name
}
