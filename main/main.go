package main

import (
	"fmt"
	"ifStructCvter/structCvter"
)

type ChildStruct struct {
	Child string `nick:"childNick"`
}

type MyStruct struct {
	OrgName   string `nick:"nickName"`
	Float     float32
	String    string
	ChildStru ChildStruct
	Hiddden   string `hide:"true"`
	private   string
}

func init() {
	structCvter.SetNickKey("nick")
	structCvter.SetFirstLetterLowerCase()
	structCvter.SetHideKey("hide")
}

func main() {
	par := MyStruct{"OrgValue", 9.99, "I am a string", ChildStruct{"child struct"}, "Hiddden value", "private value"}
	_map, err := structCvter.ToMap(par)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("map:", _map)
	}

	_json, err2 := structCvter.ToJson(par)
	if err2 != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("json:", _json)
	}
}
