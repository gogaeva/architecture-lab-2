package lab2

import (
	"strings"
	"strconv"
	"container/list"
	"errors"
)

var priorities = map [string] uint8 {
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

type chunk struct {
	exp []string
	priority uint8
}

func PostfixToInfix(postfixExp string) (string, error) {
	stack := list.New()
	stack.PushBack(chunk{[]string{""}, 0})
	symbols := strings.Split(postfixExp, " ")
	for i := 0; i < len(symbols); i++ {
		symbol := symbols[i]
		if symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/" || symbol == "^" {
			rightChunk := stack.Remove(stack.Back()).(chunk)
			if rightChunk.priority <  priorities[symbol] && rightChunk.priority != 0 {
				tmp := []string{"("}
				rightChunk.exp = append(tmp, rightChunk.exp...)
				rightChunk.exp = append(rightChunk.exp, ")")
			}
			leftChunk := stack.Remove(stack.Back()).(chunk)
			if leftChunk.priority < priorities[symbol] && leftChunk.priority != 0 {
				tmp := []string{"("}
				leftChunk.exp = append(tmp, leftChunk.exp...)
				leftChunk.exp = append(leftChunk.exp, ")")
			}
			newChunk := chunk{leftChunk.exp, priorities[symbol]}
			newChunk.exp = append(newChunk.exp, symbol)
			newChunk.exp = append(newChunk.exp, rightChunk.exp...)
			stack.PushBack(newChunk)
		} else if symbol == "" {
			continue
		} else {
			symbol, err := strconv.ParseFloat(symbol, 64)
			if err == nil {
				symbol := strconv.FormatFloat(symbol, 'G', -1, 64)
				stack.PushBack(chunk{[]string{symbol}, 0})
			} else {
				return "", errors.New("Error: invalid symbol")
			}
		}
	}
	infixExp := stack.Back().Value.(chunk).exp
	return strings.Join(infixExp, ""), nil
}
