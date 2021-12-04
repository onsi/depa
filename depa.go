package depa

import "github.com/onsi/depalpha/beta"

const A_VERSION = "1.1.0"

func A() string {
	return "A" + beta.Beta()
}
