package main

import (
	"fmt"
)

func main() {
	op := ""
	loop := true

	balance := 10000.0
	money := 0.0
	note := ""
	details := fmt.Sprintf("%-10s%-10s%-10s%-10s", "收支", "账户金额", "收支金额", "说明")
	flag := false
	for {
		fmt.Println("\n\n--------------家庭收支记账软件------------\n")
		fmt.Printf("%10s", "1 收支明细")
		fmt.Printf("%10s", "2 登记收入")
		fmt.Printf("%10s", "3 登记支出")
		fmt.Printf("%10s\n\n", "4 退出软件")
		fmt.Printf("请选择(1-4): ")

		fmt.Scanln(&op)

		switch op {
		case "1":
			fmt.Println("---------------收支明细情况-------------")
			if !flag {
				fmt.Println("没有任何收支,先来一笔?")
			} else {
				fmt.Println(details)
			}

		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			flag = true
			details += fmt.Sprintf("\n%-10s%-10v%-10v%-10s", "收入", balance, money, note)
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足!")
				break
			}
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			flag = true
			details += fmt.Sprintf("\n%-10s%-10v%-10v%-10s", "支出", balance, money, note)
		case "4":
			fmt.Println("确定要退出吗? y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}

				fmt.Println("你的输入有误,请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {
			break
		}

	}
	fmt.Println("你已退出程序!")

}
