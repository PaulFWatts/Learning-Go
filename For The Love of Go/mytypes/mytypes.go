package mytypes

import "strings"

// Twice multiplies its receiver by 2 and returns
// the result

type MyInt int
type MyString string

// MyBuilder is a custom version of the `strings.Builder` type.
type MyBuilder struct {
	Contents strings.Builder
}

// StringUppercaser wraps strings.Builder.
type StringUppercaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

// Len returns the length of the string.
func (s MyString) Len() int {
	return len(s)
}

// Hello returns the string "Hello, Gophers!".
func (mb MyBuilder) Hello() MyString {
	return MyString("Hello, Gophers!")
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func Double(input *int) {
	*input *= 2
}

func (input *MyInt) Double() {
	*input *= 2
}
