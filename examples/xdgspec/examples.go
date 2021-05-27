package xdgspec

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// impl check
var (
	_ mapping.Register = (*register)(nil)
)

func NewRegister() mapping.Register {
	return new(register)
}

func (*register) Regist(m mapping.ExampleMapping) {
	m["xdg_basedirs"] = BaseDirs
	m["xdg_userdir"] = UserDir
	m["xdg_fileop"] = FileOp
}
