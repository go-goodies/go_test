package attesto

type Assertion struct {
	t      TestHarness
	target interface{}
	To     *Matcher
}

func NewExpectation(t TestHarness, target interface{}) *Assertion {
	return &Assertion{
		t:      t,
		target: target,
		To:     &Matcher{target: target, harness: t},
	}
}
