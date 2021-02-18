package account

import "fmt"

const (
	Cint byte = 0
	Cout byte  = 1
)

type  Account struct {
	Balance float64
	Pwd string
	AccountDetails
}

type AccountDetails []AccountDetail
type AccountDetail struct {
	Money float64
	Catagory byte
	Desc string
}

//存钱
func (inputAccount *Account) Des(money float64,pwd string) {

	if inputAccount.Pwd != pwd {
		fmt.Println("输入密码不对！")
		return
	}
	fmt.Println("家庭存钱....", money)
	if money <= 0 {
		fmt.Println("存钱不合法")
	}
	inputAccount.AccountDetails = append(inputAccount.AccountDetails,AccountDetail{Money: money,Catagory: Cint,Desc: "存款"})
	inputAccount.Balance += money
}

//存钱
func (inputAccount *Account) WithDraw(money float64,pwd string) {

	if inputAccount.Pwd != pwd {
		fmt.Println("输入密码不对！")
		return
	}
	fmt.Println("家庭取钱....", money)
	if money <= 0 {
		fmt.Println("存钱不合法")
		return
	}

	if money > inputAccount.Balance {
		fmt.Println("取钱不能大于余额")
		return
	}
	inputAccount.AccountDetails = append(inputAccount.AccountDetails,AccountDetail{Money: money,Catagory: Cout,Desc: "取钱"})
	inputAccount.Balance -= money
}
//查询
func (inputAccount *Account) QueryBalance(pwd string)float64 {
	if inputAccount.Pwd != pwd {
		fmt.Println("输入密码不对！")
		return -0
	}
	fmt.Println("账户余额：",inputAccount.Balance)
	return inputAccount.Balance
}
