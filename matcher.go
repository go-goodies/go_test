package attesto

import (
	"fmt"
	"strings"
	"reflect"
)

type Matcher struct {
	target  interface{}
	message string
	harness TestHarness
}

func (matcher *Matcher) Equal(value interface{}) *Matcher {
	matcher.eval(matcher.target == value, "Expected %v to equal %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeTrue() *Matcher {
	matcher.eval(matcher.target == true, "Expected %v to be true.", matcher.target)
	return matcher
}

func (matcher *Matcher) BeFalse() *Matcher {
	matcher.eval(matcher.target == false, "Expected %v to be false.", matcher.target)
	return matcher
}

func (matcher *Matcher) BeGreaterThan(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() > value, "Expected %v to be greater than %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeGreaterThanOrEqualTo(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() >= value, "Expected %v to be greater than or equal to %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeLessThan(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() < value, "Expected %v to be less than %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeLessThanOrEqualTo(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() <= value, "Expected %v to be less than or equal to %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) Contain(value interface{}) *Matcher {

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		fmt.Println("reflect.Slice")
		vals := value.([]struct{})
		if len(vals) > 0 {

			switch reflect.TypeOf(vals[0]).Kind() {
			case reflect.String:
				for _, val := range vals {
					fmt.Printf("reflect.Slice - val: %v\n", val)
					matcher.eval(
						strings.Contains(matcher.targetAsString(), toString(val)),
						"Expected %v to contain %v.", matcher.target, val,
					)
				}
			case reflect.Int:
			for _, val := range vals {
				fmt.Printf("reflect.Slice - val: %v\n", val)

//				val_int := int(val)
//
//				if matcher.targetAsInt() != val_int {
//					matcher.eval(
//						false,
//						"Expected %v to contain %v.", matcher.target, val,
//					)
//				}


			}
//			case int:
//				padStr = t
			default:
				panic("Unknown type")
			}


		}

	case reflect.Array:
		fmt.Println("reflect.Array")
	default:
		matcher.eval(
			strings.Contains(matcher.targetAsString(), toString(value)),
			"Expected %v to contain %v.", matcher.target, value,
		)
	}
	return matcher
}

func (matcher *Matcher) targetAsInt() int {
	return int(matcher.target.(int))
}

func (matcher *Matcher) targetAsString() string {
	return toString(matcher.target)
}

func toString(value interface{}) string {
	return fmt.Sprint(value)
}

func (matcher *Matcher) eval(success bool, message string, parts ...interface{}) {

	if !success {
		matcher.message = fmt.Sprintf(message, parts...)
		matcher.harness.Log(matcher.message)
		matcher.harness.FailNow()
	}
}
