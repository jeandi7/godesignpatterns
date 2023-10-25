package behavior

import (
	"strconv"
	"strings"
	"testing"
)

func TestGetInstance(t *testing.T) {

	stack := operationStack{}
	// analyse lexicale : on mange les blancs pour en faire des lexems
	lexs := strings.Split("3 4 + 2 -", " ")

	for _, operatorString := range lexs {
		if operatorString == SUM || operatorString == SUB {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Eval())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := value(val)
			stack.Push(&temp)
		}
	}

	result := int(stack.Pop().Eval())
	if result != 5 {
		t.Errorf("must return  '%d'   not '%d' ", 5, result)
	}

	t.Log(result)

}
