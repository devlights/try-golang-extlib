package logging

import (
	"github.com/devlights/try-golang-extlib/internal/examples/logging/sentrygo"
	"github.com/devlights/try-golang/pkg/mappings"
)

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
	sentrygo.NewRegister().Regist(m)
}