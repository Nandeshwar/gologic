package main

import (
	"container/list"
	"fmt"
)

func main() {
	// infix := "a+b"
	infix := "((a+(b*c))-d)" // postfix: abc*+d-

	//infix := "(a+b)*c" // postfix: ab+c*

	postFix := convertInfixToPostfix(infix)

	fmt.Println("Infix=", infix)
	fmt.Println("Postfix=", postFix)

	infix = convertPostfixToInfix(postFix)
	fmt.Println("infix from postfix=", infix)

	fmt.Println("Algorithm2: ")
	infix = "a + b"
	postfix, prefix := convertInfixToPostfixAndPrefix(infix)
	fmt.Println("postfix=", postfix)
	fmt.Println("prefix=", prefix)
}

func convertInfixToPostfix(infix string) string {
	stack := list.New()

	var postfix []rune
	for _, v := range infix {
		c := rune(v)
		switch {
		case c >= 'a' && c <= 'z':
			fallthrough
		case c >= 'A' && c <= 'Z':
			fallthrough
		case c >= '0' && c <= '9':
			postfix = append(postfix, c)

		case c == '(':
			stack.PushBack(c)

		// pop all operator until we get '(' in stack
		case c == ')':
			for {
				element := stack.Remove(stack.Back())
				letter := element.(rune)
				if letter == '(' {
					break
				}
				postfix = append(postfix, letter)
			}

		case c == '+' || c == '-' || c == '*' || c == '/':
			// when stack is empty. insert operator and proceed to next item in loop

			if stack.Len() == 0 {
				stack.PushBack(c)
				continue
			}

			// when there is already operator in stack
			// Then, normal case: new operator precdence is > than the stack's top item priority, push that in stack
			// otherwise: pop out higher priority operator and add it result
			topElement := stack.Back().Value
			top := topElement.(rune)

			topPrecedence := findPrecedence(top)
			cPrecedence := findPrecedence(c)

			if cPrecedence > topPrecedence {
				stack.PushBack(c)
			} else {
				for {

					topElement = stack.Back().Value
					top := topElement.(rune)
					topPrecedence = findPrecedence(top)
					cPrecedence = findPrecedence(c)

					if stack.Len() >= 0 && cPrecedence <= topPrecedence {
						tElement := stack.Remove(stack.Back())
						t := tElement.(rune)
						postfix = append(postfix, t)

					} else {
						stack.PushBack(c)
						break
					}
				}
			}

		}
	}

	// pop all operators
	for stack.Len() != 0 {
		cElement := stack.Remove(stack.Back())
		c := cElement.(rune)
		postfix = append(postfix, c)
	}

	return string(postfix)
}

func findPrecedence(c rune) int {
	switch c {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func convertPostfixToInfix(postfix string) string {
	stack := list.New()

	for _, v := range postfix {
		c := rune(v)

		switch {
		case c >= 'a' && c <= 'z':
			fallthrough
		case c >= 'A' && c <= 'Z':
			fallthrough
		case c >= '0' && c <= '9':
			stack.PushBack(string(c))
		case c == '+' || c == '-' || c == '*' || c == '/':
			element1 := stack.Remove(stack.Back())
			element2 := stack.Remove(stack.Back())

			second := element1.(string)
			first := element2.(string)

			result := "(" + first + string(c) + second + ")"
			stack.PushBack(result)
		}
	}

	var infix string

	for stack.Len() > 0 {
		element := stack.Remove(stack.Back())
		s := element.(string)
		infix += s
	}

	return string(infix)
}

func convertInfixToPostfixAndPrefix(infix string) (string, string) {
	prefix := list.New()
	postfix := list.New()
	operators := list.New()

	for _, v := range infix {
		c := rune(v)

		switch {
		case c >= 'a' && c <= 'z':
			fallthrough

		case c >= 'A' && c <= 'Z':
			fallthrough

		case c >= 0 && c <= 9:
			prefix.PushBack(string(c) + " ")
			postfix.PushBack(string(c) + " ")

		case c == '(':
			operators.PushBack(string(c))

		// when operator comes
		// pop operator from stack until it reaches bracket (
		// if stack is empty - insert newly coming operator
		// if top operator is (    -> then insert newly appeared operator and out
		// else
		// keep poping top operator  and process until it is < than newly coming operator
		case c == '+' || c == '-' || c == '*' || c == '/':
			if operators.Len() == 0 {
				operators.PushBack(string(c))
				continue
			}

			opElement := operators.Back().Value
			op := opElement.(string)

			if op == "(" {
				operators.PushBack(string(c))
				continue
			}

			cPrecedence := findPrecedence(c)
			for operators.Len() > 0 {
				topOperatorElement := operators.Back().Value
				topOperator := topOperatorElement.(string)
				topOperatorPrecedence := findPrecedence(rune(topOperator[0]))
				
				if op == "(" || topOperatorPrecedence < cPrecedence {
					break
				}

				element2 := postfix.Remove(postfix.Back())
				element1 := postfix.Remove(postfix.Back())

				secondPost := element2.(string)
				firstPost := element1.(string)

				element2 = prefix.Remove(prefix.Back())
				element1 = prefix.Remove(prefix.Back())

				secondPre := element2.(string)
				firstPre := element1.(string)

				opElement := operators.Remove(operators.Back())
				op := opElement.(string)

				postfix.PushBack(firstPost + secondPost + op)
				prefix.PushBack(firstPre + secondPre + op)

			}

			operators.PushBack(string(c))
		case c == ')':
			opElement := operators.Back().Value
			op := opElement.(string)
			for op != "(" {

				element2 := postfix.Remove(postfix.Back())
				element1 := postfix.Remove(postfix.Back())

				secondPost := element2.(string)
				firstPost := element1.(string)

				element2 = prefix.Remove(prefix.Back())
				element1 = prefix.Remove(prefix.Back())

				secondPre := element2.(string)
				firstPre := element1.(string)

				postfix.PushBack(firstPost + secondPost + op)
				prefix.PushBack(firstPre + secondPre + op)

				operators.Remove(operators.Back())

				opElement = operators.Back().Value
				op = opElement.(string)

			}
			operators.Remove(operators.Back())
		}

	}

	var prefixResult string
	var postfixResult string

	for operators.Len() > 0 {
		opElement := operators.Remove(operators.Back())
		op := opElement.(string)

		element2 := postfix.Remove(postfix.Back())
		element1 := postfix.Remove(postfix.Back())

		secondPost := element2.(string)
		firstPost := element1.(string)

		element2 = prefix.Remove(prefix.Back())
		element1 = prefix.Remove(prefix.Back())

		secondPre := element2.(string)
		firstPre := element1.(string)

		postfixResult += firstPost + secondPost + op
		prefixResult += op + firstPre + secondPre

	}

	if postfix.Len() > 0 {

		element := postfix.Remove(postfix.Back())

		post := element.(string)

		element = prefix.Remove(prefix.Back())

		pre := element.(string)

		postfixResult += post
		prefixResult += pre
	}

	return postfixResult, prefixResult

}
