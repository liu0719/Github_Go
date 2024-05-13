// 2022/12/30,12:43
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1 len()长度
	str := "hello男" //utf-8中汉字占三个字节
	fmt.Println(len(str))

	//	2 有中文防止乱码要转成rune
	str2 := "rrrjsjsl北京"
	a := []rune(str2)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c\n", a[i])
	}
	fmt.Println("3____________________")
	//	3  字符串转整数 n,err:=strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(n)
	}

	fmt.Println("4--------------------")
	//4 整数转字符串  strconv.Itoa
	b := strconv.Itoa(123)
	fmt.Printf("%T\n", b)
	fmt.Println(b)

	//	Atoi  a代表字符串 to  i是int
	//Itoa    int to  a(字符串)
	fmt.Println("5------------------")
	//5 字符串转byte
	bytes := []byte("hello Go")
	fmt.Printf("%T %c\n", bytes, bytes)

	fmt.Println("6------------------")
	//6 byte转字符串
	str3 := []byte{77, 99, 88}
	fmt.Printf("%T %v\n", str3, str3)

	//7 十进制转2,8,16,进制
	fmt.Println("7-----------------------")
	n1 := strconv.FormatInt(123, 2)
	//strconv.FormatInt(要修改的数,改成的进制(2-36进制))
	fmt.Println(n1)

	//8 查找子串是否在指定字符串中
	fmt.Println("8-----------------------")
	//有则返回true 无则返回false
	fmt.Println(strings.Contains("aaabbccc", "aaj"))

	//9 查找一个字符串中有几个指定的子串 string.Count(被查找字符串，子串)
	fmt.Println("9---------------------------")
	c := strings.Count("fjtikdnffffffffdkejmdskds", "f")
	fmt.Println(c)

	//10  不区分大小写的字符串比较 strings.EqualFold()
	fmt.Println("10-----------------------")
	//区分大小写的比较就用等号
	fmt.Println("abc" == "abc", "AbC" == "aBc") //true false
	fmt.Println(strings.EqualFold("AbC", "aBc"))

	//11 返回子串在字符串第一次出现的index值，没有则返回-1
	fmt.Println("11--------------------------")
	fmt.Println(strings.Index("abc__-dfesakdjhfkdk", "sa")) //9
	fmt.Println(strings.Index("aaaaaaa", "b"))              //-1

	//12 返回子串在字符串最后一次出现的位置，没有则返回-1 strings.LastIndex()
	fmt.Println("12--------------------------")
	fmt.Println(strings.LastIndex("abcdeabcdabcabcabc", "abc"))

	//13 将指定的字符串替换strings.Replace(原字符串（s）,要替换的子串(old),替换成的新串（new),n(可以指定替换几个)-1表示全部替换)
	//替换后返回新串，原字符串不会变
	fmt.Println("13-----------------------------")
	fmt.Println(strings.Replace("go go hello world", "go", "go语言", 1))

	//14 按照指定的字符分割字符串strings.Split()
	fmt.Println("14-----------------------------")
	fmt.Println(strings.Split("hehdkk  jifwfijw  weidjiow", "j"))

	//15 字符串大小写转换
	fmt.Println("15------------------------------")
	fmt.Println(strings.ToLower("eeWWQSX"))
	fmt.Println(strings.ToUpper("efriefkeerju"))

	//16 去掉字符串两边的空格strings.trimSpace()
	fmt.Println("16----------------------------")
	fmt.Println(strings.TrimSpace(" c          d   "))

	//17 指定去掉两边的字符 Trim
	fmt.Println("17------------------------------")
	fmt.Println(strings.Trim("v vhello  r !", " !v "))
	//18  strings.TrimLeft()去掉左边
	//19  strings.TrimRight()去掉右边

	//20 判断字符串是否是以指定的字符串开头的 ，是则TRUE 不是则false
	fmt.Println("20------------------------------")
	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https://"))
	//21判断字符串是否是以指定的字符串结尾·的 ，是则TRUE 不是则false
	//	strings.HasSuffix()
}
