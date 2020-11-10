package util

import "fmt"

type CustInterface struct {
}

func (CI *CustInterface) MainInterface() {
	cs := CustServ{}
	cs.initCust()

	for {
		fmt.Println("--------------------用户信息管理软件--------------------")
		fmt.Println("1.添加用户")
		fmt.Println("2.修改用户信息")
		fmt.Println("3.查看用户列表")
		fmt.Println("4.删除用户")
		fmt.Println("0.退出软件")
		var option int
		_, _ = fmt.Scanln(&option)
		if option == 0 {
			break
		}
		switch option {
		case 1:
			cs.addCust()
		case 2:
			cs.modInfor()
		case 3:
			cs.showList()
		case 4:
			cs.deleteCust()
		default:
			fmt.Println("输入的信息有误，请重新输入")
		}
	}
}


