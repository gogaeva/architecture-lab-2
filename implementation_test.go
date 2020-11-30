package lab2

import (
	"fmt"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix_Assert (t *testing.T) {
    cases := map [string] string {
      "4 22 - 3 * 5 +" : "(4 - 22) * 3 + 5\n",
      "8 2 5 * + 1 3 2 * + 4 - /" : "(8 + 2 * 5) / (1 + 3 * 2 - 4)\n",
      "4 22 - 3 * g +" : "Algorithm error: invalid symbol",
      "8 2 5 * + 1 3 2 * +  - 4 - /" : "Algorithm error: invalid expression",
      "14.88 13.37 2.28 * + 15.3 1 ^ -" : "14.88 + 13.37 * 2.28 - 15.3 ^ 1\n",
      "222.2 52.2 - 13.67 45.2 * 2 / -" : "222.2 - 52.2 - 13.67 * 45.2 / 2\n",
      
    }
  for input, want := range cases {
    res, err := PostfixToInfix(input)
    if err == nil {
      assert.Equal(t, want, res)
    } else {
      assert.EqualError(t, err, want)
    }
  }
}

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
			want: "(4 - 22) * 3 + 5\n",
			err: nil,
		},
		{
			name: "complicated",
			arg: "8 2 5 * + 1 3 2 * + 4 - /",
			want: "(8 + 2 * 5) / (1 + 3 * 2 - 4)\n",
			err: nil,
		},
		{
			name: "invalid symbols",
			arg: "4 22 - 3 * g +",
			want: "",
			err: errors.New("Algorithm error: invalid symbol"),
		},
		{
			name: "invalid expression",
			arg: "8 2 5 * + 1 3 2 * +  - 4 - /",
			want: "",
			err: errors.New("Algorithm error: invalid expression"),
		},
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
	samples := []string{
		"4 22 - 3 * 5 +",
		"8 2 5 * + 1 3 2 * + 4 - /",
	}
	for _, sample := range samples {
		res, _ := PostfixToInfix(sample)
		fmt.Print(res)
	}
	// Output:
	// (4 - 22) * 3 + 5
	// (8 + 2 * 5) / (1 + 3 * 2 - 4)
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
