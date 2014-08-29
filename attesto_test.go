package attesto

import (
	"testing"
)

type Sample struct {
	data string
}

func TestQuiz(t *testing.T) {
	test := Test(t)

	sample1 := new(Sample)
	sample1.data = "test"

	sample2 := new(Sample)
	sample2.data = "test"

	test.Attest(true).To.BeTrue()
	test.Attest(!true).To.BeFalse()
	test.Attest(1).To.Equal(1)
	test.Attest("1").To.Equal("1")
	test.Attest(1.0).To.Equal(1.00)
	test.Attest(sample1.data).To.Equal(sample2.data)
	test.Attest(1).To.BeLessThan(2)
	test.Attest(1).To.BeLessThanOrEqualTo(1)
	test.Attest(2).To.BeGreaterThan(1)
	test.Attest(1).To.BeGreaterThanOrEqualTo(1)
	test.Attest([]string{"bar", "baz"}).To.Contain("bar")

//	slice1 := []int{0, 1, 2}
//	test.Attest(slice1).To.Contain([]int{0, 1, 2})

//	slice2 := []string{"0", "1", "2", "3"}
//	test.Attest(slice2).To.Contain([]string{"0", "1", "2"})

}
