package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "os"
	// "system"
)

var (
	result, input float64
	opt           byte = '+'
)

func cal() {
	switch opt {
	case '-':
		result -= input
	case '*':
		result *= input
	case '/':
		result /= input
	default:
		result += input
	}
}

func renumberInput(s string) (ok bool) {
	if s == "0" && opt == '/' {
		fmt.Println("-----Error Input")
		return
	}

	if f, err := strconv.ParseFloat(s, 32); err == nil {
		input = f
		cal()
		ok = true
	} else {
		fmt.Println("-----Error Input")
	}
	return
}

func main() {
	var in string
	var line int
	var tmp byte
	var ok bool

	// exit := true
	// for exit {
	for {
		line, _ = fmt.Scanln(&in)
		if line == 0 {
			continue
		}
		if in == "q" {
			break
		}
		if in == "c" {
			result = 0
			continue
		}

		switch in[0] {
		case '+':
			tmp = opt
			opt = '+'
			if len(in) > 1 {
				ok = renumberInput(in[1:])
				if !ok {
					opt = tmp
				}
			}
		case '-':
			tmp = opt
			opt = '-'
			if len(in) > 1 {
				ok = renumberInput(in[1:])
				if !ok {
					opt = tmp
				}
			}
		case '*':
			tmp = opt
			opt = '*'
			if len(in) > 1 {
				ok = renumberInput(in[1:])
				if !ok {
					opt = tmp
				}
			}
		case '/':
			tmp = opt
			opt = '/'
			if len(in) > 1 {
				ok = renumberInput(in[1:])
				if !ok {
					opt = tmp
				}
			}
		case '=':
			fmt.Printf("----------------\n%v\n", result)
			// exit = false
		default:
			renumberInput(in)
		}
	}
}
