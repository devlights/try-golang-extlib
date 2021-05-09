package xdgspec

import "github.com/devlights/try-golang/pkg/mappings"

type (
	register struct{}
)

// impl check
var (
	_ mappings.Register = (*register)(nil)
)

func NewRegister() mappings.Register {
	return new(register)
}

func (*register) Regist(m mappings.ExampleMapping) {
	m["xdg_basedirs"] = BaseDirs
	m["xdg_userdir"] = UserDir
	m["xdg_fileop"] = FileOp
}