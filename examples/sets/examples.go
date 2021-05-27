package sets

import (
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mapping.Register {
	r := new(register)
	return r
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["set_newset"] = NewSet
	m["set_union"] = Union
	m["set_difference"] = Difference
	m["set_intersect"] = Intersect
	m["set_symmetricdiff"] = SymmetricDiff
	m["set_usingmap"] = UsingMap
}
