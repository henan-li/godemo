package utils

import "fmt"

type FamilyAccount struct {
	key         string
	loop        bool
	balance     float64
	money       float64
	note        string
	details     string
	trans_count float32
}

func NewFamilyAccount() *FamilyAccount {

	obj := &FamilyAccount{
		key:         "",
		loop:        false,
		balance:     10000.00,
		money:       0.00,
		note:        "",
		details:     "\nin\t\tout\t\tbalance\t\tnote\t\ttransaction numbers\n",
		trans_count: 0,
	}
	return obj

}

func (this *FamilyAccount) MainMenu() {

	this.loop = true

	for {
		fmt.Println("money track")
		fmt.Println("1.details")
		fmt.Println("2.money in")
		fmt.Println("3.money out")
		fmt.Println("4.exit")
		fmt.Print("select 1 to 4\n")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.transIn()
		case "3":
			this.transOut()
		case "4":
			this.doExit()
		default:
			fmt.Println("choose number between 1 and 4")
		}

		if this.loop {
			break
		}
	}
}

func (this *FamilyAccount) transIn() {
	fmt.Println("money in")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("money in note")
	fmt.Scanln(&this.note)
	this.trans_count += 1
	this.details += fmt.Sprintf("\n%v\t%v\t%v\t%v\t\t%v", this.money, "0.00", this.balance, this.note, this.trans_count)
	fmt.Println(this.details)
}

func (this *FamilyAccount) transOut() {
	fmt.Println("money out")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		//
		//this.key = ""
		//this.balance = 10000.00
		//this.money = 0.00
		//this.note = ""
		//this.details = "\nin\t\tout\t\tbalance\t\tnote\t\ttransaction numbers\n"
		//this.trans_count = 0

		fmt.Println("balance is limited, start over")
		fmt.Println("money track")
		fmt.Println("1.details")
		fmt.Println("2.money in")
		fmt.Println("3.money out")
		fmt.Println("4.exit")
		fmt.Print("select 1 to 4\n")

		fmt.Scanln(&this.key)
	}

	if this.loop == true {
		this.balance -= this.money
		fmt.Println("money out note")
		fmt.Scanln(&this.note)
		this.trans_count += 1
		this.details += fmt.Sprintf("\n%v\t%v\t%v\t%v\t\t%v", "0.00", this.money, this.balance, this.note, this.trans_count)
		fmt.Println(this.details)
	}

}

func (this *FamilyAccount) doExit() {
	fmt.Println("exit for real? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("input your choice again? y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) showDetails() {
	if this.trans_count == 0 {
		fmt.Println("no records, pls make new in or out record first")
	} else {
		fmt.Println(this.details)
	}
}
