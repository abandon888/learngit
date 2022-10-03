package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "Hello南昌"
	r := []rune(str1)
	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("字符串=%c\n", r[i])
	// }
	//使用遍历的方法输出
	for i, b := range r {
		fmt.Printf("i=%v,b=%c\n", i, b)
	}
	//把字符串转数字
	n, err := strconv.Atoi("e")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}
	// 整数转字符串
	str := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T", str, str)
	//字符串转[]byte（AS编码）:
	var bytes = []byte("hello,golang")
	fmt.Printf("str=%v", bytes)
	//[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)
	//查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "sea")
	fmt.Printf("b=%v\n", b)
	//使用new创建的是地址，指针（会分配内存）
	num2 := new(int)
	*num2 = 400 //如此赋值
	fmt.Printf("num2类型%T，num2地址%d,num2的值%d，指向的数%d", num2, &num2, num2, *num2)

}
