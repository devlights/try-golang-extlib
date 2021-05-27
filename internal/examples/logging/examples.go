package logging

import (
	"github.com/devlights/try-golang-extlib/internal/examples/logging/sentrygo"
	"github.com/devlights/try-golang/mapping"
)

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
	sentrygo.NewRegister().Regist(m)
}