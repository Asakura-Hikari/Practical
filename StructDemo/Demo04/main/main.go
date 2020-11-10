package main

import "fmt"

/*
随手写的银行存储系统, 虽然全部都存在main包里真的很奇怪。。。
*/

type Account struct {
	Acc   int
	Pwd   string
	Money float64
}

func (acc *Account) deposit() {
	var dep float64
	fmt.Println("输入你要存入的金额")
	_, _ = fmt.Scanln(&dep)
	acc.Money = acc.Money + dep
	fmt.Printf("此账号余额为: %.2f\n", acc.Money)
}

func (acc *Account) withdraw() {
	var wtd float64
	fmt.Println("输入你要取出的金额")
	_, _ = fmt.Scanln(&wtd)
	acc.Money = acc.Money - wtd
	fmt.Printf("此账号余额为: %.2f\n", acc.Money)
}

func (acc *Account) check() {
	fmt.Printf("此账号剩余金额: %.2f\n", acc.Money)
}

func login(account []Account) (bool, Account) {
	var ac int
	var pwd string
	var log bool
	fmt.Println("请输入你的账号")
	_, _ = fmt.Scanln(&ac)
	fmt.Println("请输入你的密码")
	_, _ = fmt.Scanln(&pwd)

	for _, v := range account {
		if ac == v.Acc && pwd == v.Pwd {
			fmt.Println("登陆成功")
			log = true
			return log, v
		}
	}
	fmt.Println("用户名或者密码不正确")
	return log, account[0]
}

func regist(account []Account) ([]Account,bool) {
	var Acc int
	var Pwd string
	var Money float64
	var re bool
	fmt.Println("请输入你的账号")
	_, err1 := fmt.Scanln(&Acc)
	if err1 != nil {
		fmt.Println("你输入的账号有误")
		return account, re
	}
	fmt.Println("请输入你的密码")
	_, _ = fmt.Scanln(&Pwd)
	fmt.Println("请输入你存入的金额")
	_, err := fmt.Scanln(&Money)
	if err != nil {
		fmt.Println("你输入的金额有误")
	}
	account = append(account, Account{Acc, Pwd, Money})
	return account, re
}

func (acc *Account) control() {
	var input int

	level1:
	for {
		fmt.Println("————————————输入以下数字进行操作————————————")
		fmt.Println("1.查询")
		fmt.Println("2.存款")
		fmt.Println("3.取款")
		fmt.Println("0.返回上级")
		_, _ = fmt.Scanln(&input)
		switch input {
		case 1:
			acc.check()
		case 2:
			acc.deposit()
		case 3:
			acc.withdraw()
		case 0:
			fmt.Println("0.返回上级成功")
			break level1
		}
		fmt.Println()
	}
}

func main() {
	var acc = make([]Account, 1)
	var control int
	acc[0] = Account{000001, "55555", 30000}

	level1:
	for {
		fmt.Println("————————————输入以下数字进行操作————————————")
		fmt.Println("1.登陆")
		fmt.Println("2.注册")
		fmt.Println("0.退出")
		_, err := fmt.Scanln(&control)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch control {
		case 1:
			log, account := login(acc)
			if log {
				account.control()
			} else {
				break
			}
		case 2:
			acco,re := regist(acc)
			if re == true {
				fmt.Println("注册成功！")
				acc = acco
			}else {
				fmt.Println("注册失败...")
				break
			}
		case 0:
			fmt.Println("退出成功")
			break level1
		}
		fmt.Println()
	}
}
