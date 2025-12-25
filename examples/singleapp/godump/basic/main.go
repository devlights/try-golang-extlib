package main

import "github.com/goforj/godump"

type (
	Like struct {
		Id    int
		Value string
	}
	Data struct {
		Name  string
		Age   uint
		Likes []Like
	}
)

func NewData(name string, age uint) *Data {
	d := Data{
		Name: name,
		Age:  age,
	}

	d.Likes = make([]Like, 0)

	return &d
}

func (me *Data) AddLike(v string) {
	me.Likes = append(me.Likes, Like{Id: len(me.Likes) + 1, Value: v})
}

func main() {
	d1 := NewData("FOO", 33)
	d1.AddLike("music")
	d1.AddLike("cycling")

	// ダンプ
	godump.Dump(d1)

	// もう一つデータを生成して比較
	d2 := NewData("foo", 44)
	d2.AddLike("music")
	d2.AddLike("cycling")
	d2.AddLike("programming")

	// 比較
	godump.Diff(d1, d2)
}

/*
	$ task
	<#dump // main.go:38
	#*main.Data {
	+Name  => "FOO" #string
	+Age   => 33 #uint
	+Likes => #[]main.Like [
		0 => #main.Like {
		+Id    => 1 #int
		+Value => "music" #string
		}
		1 => #main.Like {
		+Id    => 2 #int
		+Value => "cycling" #string
		}
	]
	}
	<#diff // main.go:47
	#*main.Data {
	-   +Name  => "FOO" #string
	-   +Age   => 33 #uint
	+   +Name  => "foo" #string
	+   +Age   => 44 #uint
		+Likes => #[]main.Like [
		0 => #main.Like {
			+Id    => 1 #int
			+Value => "music" #string
		}
		1 => #main.Like {
			+Id    => 2 #int
			+Value => "cycling" #string
		}
	+     2 => #main.Like {
	+       +Id    => 3 #int
	+       +Value => "programming" #string
	+     }
		]
	}
*/
