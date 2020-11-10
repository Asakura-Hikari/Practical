package util

import (
	"fmt"
)

type CustServ struct {
	CustList []Customer
	custNum  int
}

func (CS *CustServ) initCust(){
	CS.custNum = 1
	Cust := Customer{Id: CS.custNum, Name: "阿卡丽", Gander: "女", Age: 20, Phone: "17254727364",
		Email: "Akaili@gmail.com"}
	CS.CustList = append(CS.CustList, Cust)
}

func (CS *CustServ) addCust() {
	name := ""
	gander := ""
	age := 0
	phone := ""
	email := ""

	fmt.Println("请输入新用户名")
	_, _ = fmt.Scanln(&name)

	for {
		fmt.Println("请输入新用户性别（不可更改）：")
		_, _ = fmt.Scanln(&gander)
		if gander == "男" || gander == "女" {
			break
		} else {
			fmt.Println("你输入的信息有误，请重新输入")
		}
	}

	for {
		fmt.Println("请输入新用户年龄（不可更改）：")
		_, _ = fmt.Scanln(&age)
		if age > 0 || age < 100 {
			break
		} else {
			fmt.Println("你输入的年龄有误，请重新输入")
		}
	}

	for {
		fmt.Println("请输入新用户电话：")
		_, _ = fmt.Scanln(&phone)
		if len(phone) == 11 {
			break
		} else {
			fmt.Println("你输入的电话号码有误，请重新输入")
		}
	}

	fmt.Println("请输入新用户邮箱：")
	_, _ = fmt.Scanln(&email)

	CS.custNum++

	cust := Customer{Id: CS.custNum, Name: name, Gander: gander, Age: age, Phone: phone, Email: email}
	CS.CustList = append(CS.CustList, cust)
	fmt.Println("添加成功")
}

func (CS *CustServ) showList() {
	fmt.Printf("id\t姓名\t性别\t年龄\t电话\t\t邮箱\n")
	for _, v := range CS.CustList {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", v.Id, v.Name, v.Gander, v.Age, v.Phone, v.Email)
	}
	fmt.Printf("------------------------用户列表------------------------\n\n")
}

func (CS *CustServ) deleteCust() {
	id := 0
	of := false
	for {
		fmt.Println("请输入要删除用户的ID(输入0退出)：")
		_, _ = fmt.Scanln(&id)

		if id == 0 {
			return
		}

		for i, v := range CS.CustList {
			if id == v.Id {
				of = true
				var del string
				fmt.Println("确定要删除吗？（Y/N）")
				_, _ = fmt.Scanln(&del)

				if del == "y" || del == "Y" {
					CS.CustList = append(CS.CustList[:i], CS.CustList[i+1:]...)
					fmt.Println("删除成功")
					return

				} else if del == "n" || del == "N" {
					break

				} else {
					fmt.Println("输入的信息有误，请重新输入")
				}
			}
		}
		if of == false {
			fmt.Println("删除失败，列表中没有该用户")
		}
	}
}

func (CS *CustServ) modInfor() {
	fmt.Println("请输入需要修改的用户ID")
	for {
		id := 0
		_, _ = fmt.Scanln(&id)
		for i:= 0; i< len(CS.CustList);i++ {
			if id == CS.CustList[i].Id {
				for {
					fmt.Println("--------------------请输入你要修改的内容--------------------")
					fmt.Println("1.姓名")
					fmt.Println("2.电话")
					fmt.Println("3.邮箱")
					fmt.Println("4.返回上级")
					option := 0
					_, _ = fmt.Scanln(&option)
					if option == 4 {
						return
					}
					switch option {
					case 1:
						fmt.Println("输入你想修改的名字，当前名字为", CS.CustList[i].Name)
						name := ""
						_, _ = fmt.Scanln(&name)
						CS.CustList[i].Name = name
						fmt.Println("修改成功")

					case 2:
						fmt.Println("输入你想修改的电话，当前电话为", CS.CustList[i].Phone)
						phone := ""
						_, _ = fmt.Scanln(&phone)
						CS.CustList[i].Phone = phone
						fmt.Println("修改成功")

					case 3:
						fmt.Println("输入你想修改的邮箱，当前邮箱为", CS.CustList[i].Email)
						email := ""
						_, _ = fmt.Scanln(&email)
						CS.CustList[i].Email = email
						fmt.Println("修改成功")

					default:
						fmt.Println("你输入的信息有误，请重新输入")
					}
				}
			}
		}
		fmt.Println("列表中没有该用户，请重新输入")
	}
}
