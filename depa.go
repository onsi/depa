package depa

import "github.com/onsi/depalpha/beta"

const A_VERSION = "1.3.0"

func A() string {
	return "A" + beta.Beta()
}
