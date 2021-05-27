package sentrygo

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
	m["logging_sentry_basic"] = Basic
	m["logging_sentry_bad_pattern"] = BadPattern
	m["logging_sentry_good_pattern"] = GoodPattern
}
