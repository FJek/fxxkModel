package main

import (
	"fmt"
	"fzw/fxxkModel/util"
)

func main() {
	snakeStr := util.CamelStringReg("Japan_Canada_Australia_Hello_World")
	fmt.Println(snakeStr)
}
