package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main()  {
	//将FF转换位16进制的整型
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	//将12转换为10进制的int型整数
	fmt.Println(strconv.ParseUint("12", 10, 0))
	// 解析一个表示浮点数的字符串并返回其值。
	v:="4.23254578"
	//精度不够时
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	//精度足够时
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	//返回子串在字符串s中第一次出现的索引值，不在的话返回-1.
	str:="我是中国人"
	fmt.Println(strings.Index(str, "我"))
	fmt.Println(strings.Index(str, "你"))

	//查找子串是否在指定的字符串中
	fmt.Println(strings.Contains("我们都是中国人", "中国人"))
	//空的字串可以存在于空的字符串中
	fmt.Println(strings.Contains("", ""))
}

