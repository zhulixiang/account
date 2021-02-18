package menu

import (
	"familyAccount/account"
	"fmt"
)

type Menu struct {
}


func NewMenu() *Menu{
	return &Menu{}
}


func (menu *Menu) ShowMenu() {
	fmt.Println("家庭财务管理，菜单功能如下：")
	fmt.Println("1.收支明细")
	fmt.Println("2.登记收入")
	fmt.Println("3.登记支出")
	fmt.Println("4.退出")
	menu.OperateAccount()
}

func (menu *Menu) OperateAccount() {
	var pwd string = "666666"
	var account = account.Account{Balance: 0,Pwd: pwd}
	for {
		fmt.Println("请输入操作的功能编号：")
		var catagory int
		fmt.Scanln(&catagory)
		var isContinue bool = true
		switch catagory {
		case 1:
			if len(account.AccountDetails) <= 0 {
				fmt.Println("没有收入登记明细")
			}else {
				fmt.Printf("收支\t账户金额\t收支金额\t说明\n")
				for _ ,val := range account.AccountDetails{
					catagory := ""
					if val.Catagory == 1 {
						catagory = "支出"
					}
					if val.Catagory == 0 {
						catagory = "收入"
					}
					fmt.Printf("%v\t%v\t\t\t%v\t%v\n",catagory,account.Balance,val.Money,val.Desc)
				}
			}
		case 2:
			account.Des(100,pwd)
		case 3:
			account.WithDraw(10,pwd)
		case 4:
			fmt.Println("退出")
			isContinue = false
		default:
			fmt.Println("没有该功能，退出")
			isContinue = false
		}
		if !isContinue {
			break
		}
	}
}
