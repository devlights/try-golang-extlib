package examples

import (
	"github.com/devlights/try-golang-extlib/internal/examples/generate"
	"github.com/devlights/try-golang-extlib/internal/examples/gocmp"
	"github.com/devlights/try-golang-extlib/internal/examples/logging"
	"github.com/devlights/try-golang-extlib/internal/examples/sets"
	"github.com/devlights/try-golang-extlib/internal/examples/xdgspec"
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() mapping.Register {
	r := new(register)
	return r
}

// Regist は、advanced パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	generate.NewRegister().Regist(m)
	gocmp.NewRegister().Regist(m)
	sets.NewRegister().Regist(m)
	logging.NewRegister().Regist(m)
	xdgspec.NewRegister().Regist(m)
}
