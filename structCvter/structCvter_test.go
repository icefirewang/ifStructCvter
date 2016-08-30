package structCvter

import (
	"fmt"
	"testing"
)

type ChildStruct struct {
	Child string `nick:"childNick"`
}

type MyStruct struct {
	Abc   string `nick:"NickAbc"`
	Abcd  float32
	Ccc   string
	Child ChildStruct
	//	Public  string
	private string
}

type MapStruct struct {
	Map map[string]interface{} `nick:"Im map"`
}

func init() {
	SetNickKey("nick")
	SetFirstLetterLowerCase()
	SetHideKey("hide")
}

func Test_ModelConvertToJSon(t *testing.T) {
	println("Test_ModelConvertToJSon")
	par := MyStruct{"abcd", 1.5, "????", ChildStruct{"child"}, "public"}
	//	par := MyStruct{"a", 1, "c"}
	str, _err := ToJson(par)
	println(_err)
	println("json", str)
}

func Test_ToMap(t *testing.T) {

	println("Test_ToMap")
	par := MyStruct{"abcd", 1.5, "????", ChildStruct{"child"}, "public"}
	//	par := MyStruct{"a", 1, "c"}
	ret, err := ToMap(par)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}

func Test_StructWithMap(t *testing.T) {
	println("Test_StructWithMap")
	m := make(map[string]interface{})
	mv := MyStruct{"abcd", 1.5, "????", ChildStruct{"child"}, "public"}
	m["a"] = mv
	par := MapStruct{m}
	str, _err := ToJson(par)
	println(_err)
	println("map json", str)
}

type ComplexStruct struct {
	Map    map[string]interface{} `nick:"I am a map"`
	Abc    string                 `nick:"NickAbc"`
	Abcd   float32
	CccDdd string

	ChildStruct MyStruct
	sss         string
	//	small       string
}

//
func Test_ComplexStruct(t *testing.T) {
	fmt.Println("Test_ComplexStruct")
	m := make(map[string]interface{})
	m["a"] = MyStruct{"abcd", 1.5, "????", ChildStruct{"child"}, "public"}
	m["b"] = "string"
	m["c"] = 123
	m["d"] = 1.111
	child := MyStruct{"abcd", 1.5, "????", ChildStruct{"child"}, "public"}
	stru := ComplexStruct{m, "abc", 9.99, "string aaa", child, "small string"}
	//stru1 := ComplexStruct{child, "ss"}
	_map, err := ToMap(stru)
	jsonStr, err2 := ToJson(stru)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("map string :", _map)
	}

	if err2 != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("json string :", jsonStr)
	}
}
