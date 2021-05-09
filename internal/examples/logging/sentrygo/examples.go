package sentrygo

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
	m["logging_sentry_basic"] = Basic
	m["logging_sentry_bad_pattern"] = BadPattern
	m["logging_sentry_good_pattern"] = GoodPattern
}