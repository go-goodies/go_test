package attesto

import (
	"log"
	"testing"
)

type defaultHarness struct {
	t *testing.T
}

func (harness defaultHarness) FailNow() {
	harness.t.FailNow()
}

func (harness defaultHarness) Log(line string) {
	log.Printf(line)
}

func (harness defaultHarness) Attest(target interface{}) *Assertion {
	return NewExpectation(harness, target)
}
