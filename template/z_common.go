package template

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

func parse() {
	//将字符串转换成10进制整型, output : 123
	s1, _ := strconv.Atoi("123")
	fmt.Println(s1)
	//解析整数, base:进制, bitSize:范围限制, output : int64, 155
	s2, _ := strconv.ParseInt("010011011", 2, 64)
	fmt.Printf("%T, %v \n", s2, s2)

	//解析浮点数, output : float64, 3.1415
	s3, _ := strconv.ParseFloat("3.1415",64)
	fmt.Printf("%T, %v \n", s3, s3)

	//将整型转换成10进制字符串, output : "123"
	s4 := strconv.Itoa(123)
	fmt.Println(s4)
	//将整型转换成 base 进制字符串, output : "1111011"
	s5 := strconv.FormatInt(123, 2)
	fmt.Println(s5)

	//无符号整数二进制长度, output : 4
	s6 := bits.Len(uint(10))
	fmt.Println(s6)
}

func strOperation() {
	//以 sep 分割字符串, output : [a b c d]
	str := strings.Split("a,b,c,d", ",")
	fmt.Println(str)
	//将字符数组添加字符变为字符串, output : "i,am,iron,man"
	str1 := strings.Join([]string{"i", "am", "iron", "man"}, ",")
	fmt.Println(str1)

	//将字符串中的字符替换为新字符, output : "i am iron man"
	str2 := strings.ReplaceAll("i,am,iron,man", ",", " ")
	fmt.Println(str2)
	//统计字符串字符的数量, output : 2
	s := strings.Count("i,am,iron,man", "a")
	fmt.Println(s)

	//删除字符串所有前导字符, output : 123123
	str3 := strings.TrimLeft("00123123", "0")
	fmt.Println(str3)
	//删除字符串前缀 prefix, output : 0123123
	str4 := strings.TrimPrefix("00123123", "0")
	fmt.Println(str4)
	//删除字符串所有前导和后缀字符, output : 123123
	str5 := strings.Trim("0012312300", "0")
	fmt.Println(str5)
	//检查字符串是否有前缀 prefix, output : true
	b := strings.HasPrefix("00123123", "0")
	fmt.Println(b)
}

func min(a, b int) int {if a < b {return a}; return b}
func max(a, b int) int {if a > b {return a}; return b}
func abs(a int) int {if a < 0 {return -a}; return a}
