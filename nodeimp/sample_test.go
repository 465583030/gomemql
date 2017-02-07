package gomemql

import (
	"testing"
)

type tableDef struct {
	Id    int32
	Level int32
	Name  string
}

var tabData = []*tableDef{
	&tableDef{Id: 6, Level: 20, Name: "kitty"},
	&tableDef{Id: 1, Level: 50, Name: "hello"},
	&tableDef{Id: 4, Level: 20, Name: "kitty"},
	&tableDef{Id: 5, Level: 10, Name: "power"},
	&tableDef{Id: 3, Level: 20, Name: "hello"},
	&tableDef{Id: 2, Level: 20, Name: "kitty"},
}

func TestHelloWorld(t *testing.T) {

	tab := NewTable()

	for _, v := range tabData {
		tab.AddRecord(v.Name, v)
	}

	// 匹配Name为hello
	NewQuery(tab).Equal("hello").Result(func(v interface{}) bool {

		t.Log(v)

		return true
	})

}

func Test2ConditionWithIndex(t *testing.T) {

	tab := NewTable()

	for _, v := range tabData {
		tab.AddRecord(v.Name, v.Id, v)
	}

	// 构建第二个字段(Id), 从1~6的索引
	tab.GenIndexNotEqual(1, 1, 6)

	NewQuery(tab).Equal("kitty").NotEqual(int32(4)).Result(func(v interface{}) bool {

		t.Log(v)

		return true
	})

}