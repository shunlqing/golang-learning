package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "shunshunshunlqingshun"
	str1 := strings.Trim(str, "shun")
	str2 := strings.TrimLeft(str, "shun")
	str3 := strings.TrimPrefix(str, "shun")
	fmt.Println(str1, str2, str3, str)

	str = "shnlqing add: sdfadf SH"
	strs := strings.Fields(str)
	str4 := strings.Join(strs, "&")
	fmt.Println(strs, str4)
}