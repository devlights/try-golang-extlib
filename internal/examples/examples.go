package examples

import (
	"github.com/devlights/try-golang-extlib/internal/examples/generate"
	"github.com/devlights/try-golang/pkg/mappings"
)

type (
	register struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist は、advanced パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	generate.NewRegister().Regist(m)
}
