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
	exp string
	priority uint8
}

func popChunck (stack *list.List) (chunk, error) {
	lastElem := stack.Back()
	if lastElem == nil {
		return chunk{"", 0}, errors.New("Algorithm error: invalid expression")
	}
	return stack.Remove(lastElem).(chunk), nil
} 

func PostfixToInfix(postfixExps string) (string, error) {
	stack := list.New()
	var infixExps string
	expressions := strings.Split(postfixExps, "\n")
	if expressions[len(expressions)-1] == "" {
		expressions = expressions[:len(expressions)-1]
	}
	for _, expression := range expressions {
		symbols := strings.Split(expression, " ")
		for i := 0; i < len(symbols); i++ {
			symbol := symbols[i]
			if symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/" || symbol == "^" {
				rightChunk, err := popChunck(stack)
				if err != nil {
					return "", err
				}
				if rightChunk.priority <  priorities[symbol] && rightChunk.priority != 0 {
					rightChunk.exp = "(" + rightChunk.exp + ")"
				}
				leftChunk, err := popChunck(stack)
				if err != nil {
					return "", err
				}
				if leftChunk.priority < priorities[symbol] && leftChunk.priority != 0 {
					leftChunk.exp = "(" + leftChunk.exp + ")"
				}
				newChunk := chunk{leftChunk.exp + " " + symbol + " " + rightChunk.exp, priorities[symbol]}
				stack.PushBack(newChunk)
			} else if symbol == "" {
				continue
			} else {
				symbol, err := strconv.ParseFloat(symbol, 64)
				if err != nil {
					return "", errors.New("Algorithm error: invalid symbol")
				} else {
					symbol := strconv.FormatFloat(symbol, 'G', -1, 64)
					stack.PushBack(chunk{symbol, 0})
				}
			}
		}
		if stack.Len() > 1 {
			return "", errors.New("Algorithm error: invalid expression")
		}
		next, err := popChunck(stack)
		if err != nil {
			return "", err
		}	
		infixExps += next.exp + "\n"
	}
	return infixExps, nil
}
