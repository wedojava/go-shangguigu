package utils

import "fmt"

type FamilyAccount struct {
	op      string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
	flag    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		op:      "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: fmt.Sprintf("%-10s%-10s%-10s%-10s", "收支", "账户金额", "收支金额", "说明"),
		flag:    false,
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------收支明细情况-------------")
	if !this.flag {
		fmt.Println("没有任何收支,先来一笔?")
	} else {
		fmt.Println(this.details)
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n%-10s%-10v%-10v%-10s", "收入", this.balance, this.money, this.note)
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足!")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.flag = true
	this.details += fmt.Sprintf("\n%-10s%-10v%-10v%-10s", "支出", this.balance, this.money, this.note)
}

func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n\n--------------家庭收支记账软件------------\n")
		fmt.Printf("%10s", "1 收支明细")
		fmt.Printf("%10s", "2 登记收入")
		fmt.Printf("%10s", "3 登记支出")
		fmt.Printf("%10s\n\n", "4 退出软件")
		fmt.Printf("请选择(1-4): ")

		fmt.Scanln(&this.op)

		switch this.op {
		case "1":
			this.showDetails()

		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !this.loop {
			break
		}

	}
	fmt.Println("你已退出程序!")
}
