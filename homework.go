package main

import (
	"fmt"
	"math"
	"strconv"
)

func compute(a, b, c, d float64, sign string) complex128 {
	var res complex128

	switch sign {
	case "+":
		res = complex(a+c, b+d)
		return res
	case "-":
		res = complex(a-c, b-d)
		return res
	case "*":
		res = complex((a*c - b*d), (b*c + a*d))
		return res
	case "/":
		res = complex((a*c+b*d)/(c*c+d*d), (b*c-a*d)/(c*c+d*d))
		return res
	default:
		fmt.Println("请正确输入！")
		return 0
	}
}
func compute1(a, b float64, sign string) float64 {
	var res float64
	switch sign {
	case "求模长":
		res = math.Sqrt(a*a + b*b)
		return res
	default:
		fmt.Println("请正确输入！")
		return 0
	}
}
func toString(a float64) string {
	res := strconv.FormatFloat(a, 'b', 1, 64)
	return res
}
func main() {
	var a, b, c, d, result1 float64
	var sign, result2 string
	var result complex128
	fmt.Println("规则说明：目前仅支持运算符号+，-，*，/，tostring（仅计算第一位复数的实部），求模长（仅计算第一位复数）")
	fmt.Println("请输入第一位复数的实部与虚部")
	fmt.Scanln(&a, &b)
	fmt.Println("请输入第二位复数实部与虚部")
	fmt.Scanln(&c, &d)
	fmt.Println("请输入运算符号(+，-，*，/，tostring，求模长)")
	fmt.Scanln(&sign)
	if len(sign) == 9 {
		result1 = compute1(a, b, sign)
		fmt.Println("运算结果为", result1)
	} else if len(sign) == 8 {
		result2 = toString(a)
		fmt.Println("运算结果为", result2)
	} else {
		result = compute(a, b, c, d, sign)
		fmt.Println("运算结果为", result)
	}

}
