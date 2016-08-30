# ifStructCvter
A package to convert  go struct to  map or json.

If struct has child struct will be ok too.

===
 
一个将 go struct 转化成 map 或者 json 的包。

如果 struct 里面包含 struct，也一样 OK



## Example
```go

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
```

