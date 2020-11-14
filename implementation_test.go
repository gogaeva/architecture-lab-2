package lab2

import (
	"fmt"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix (t *testing.T) {
	res, err := PostfixToInfix("4 22 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(4-22)*3+5", res)
	}
}

// func TestConvert (t *testing.T) {
// 	res, _ := Convert("4 22 - 3 * 5 +")
// 	want := "(4-22)*3+5"
// 	assert.Equal(t, res, want, "Incorrect result")
// }

func TestTablePostfixToInfix (t *testing.T) {
	cases := []struct {
		name string
		arg string
		want string
		err error
	}{
		{
			name: "simple",
			arg: "4 22 - 3 * 5 +",
			want: "(4-22)*3+5",
			err: nil,
		},
		{
			name: "complicated",
			arg: "8 2 5 * + 1 3 2 * + 4 - /",
			want: "(8+2*5)/(1+3*2-4)",
			err: nil,
		},
		{
			name: "invalid symbols",
			arg: "4 22 - 3 * g +",
			want: "",
			err: errors.New("Error: invalid symbol"),
		},
		// {
		// 	name: "invalid expression",
		// 	arg: "8 2 5 * + 1 3 2 * +  - 4 - /",
		// 	want: "",
		// 	err: errors.New("Error: invalid expression (redundant symbols in expression)"),
		// },
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			 res, err := PostfixToInfix(testCase.arg)
			 if res != testCase.want {
			 	t.Errorf("Expercted: %s, got: %s", testCase.want, res)
			 }
			 if err != nil && err.Error() != testCase.err.Error() {
			 	t.Errorf("Expected: %s, got: %s", testCase.err.Error(), err.Error())
			 }
		})
	}
}

func ExamplePostfixToInfix () {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2+2
}

func BenchmarkPostfixToInfix (b *testing.B) {
	b.Run("short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostfixToInfix("4 22 - 3 * 5 +")
		}
	})
	b.Run("long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostfixToInfix("8 2 5 * + 1 3 2 * + 4 - /")
		}
	})
}