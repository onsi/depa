package depa_test

import (
	"testing"

	"github.com/onsi/depa"
	"github.com/onsi/depalpha/beta"
)

func TestDepA(t *testing.T) {
	if depa.A()+beta.Beta() != "ABeta" {
		t.Fatal("bloop")
	}
}
