package attesto

type TestHarness interface {
	FailNow()
	Log(string)
	Attest(interface{}) *Assertion
}
