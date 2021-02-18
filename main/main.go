package main

import (
	"familyAccount/menu"
	"fmt"
)


func main(){
	fmt.Println("家庭收支项目。。。")
	menu.NewMenu().ShowMenu()
}
