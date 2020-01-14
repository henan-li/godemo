package main

import (
	"fmt"
	"math/rand"
)
import "../controller"
import "../model"

type CustomerView struct {
	key  string
	loop bool
	customerController *controller.CustomerController
}

func (this *CustomerView) MainMenu() {
	for {
		if !this.loop {
			break
		}

		fmt.Println("user manage")
		fmt.Println("1.add")
		fmt.Println("2.edit")
		fmt.Println("3.delete")
		fmt.Println("4.show list")
		fmt.Println("5.exit")
		fmt.Print("select 1 to 5\n")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			this.edit()
		case "3":
			this.delete()
		case "4":
			this.showList()
		case "5":
			this.doExit()
		default:
			fmt.Println("choose number between 1 and 5")
		}
	}
}

func (this *CustomerView) showList() {
	customers := this.customerController.List()
	fmt.Println("-- customer list --")
	fmt.Println("id\tname\tgender\tage\tphone\temail")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Print("-- customer list end --")
}

func (this *CustomerView) add() {
	fmt.Println("-- add customer --")
	fmt.Println("name:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("gender:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("phone:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("email:")
	email := ""
	fmt.Scanln(&email)

	fmt.Println("age:")
	age := 0
	fmt.Scanln(&age)

	rand_id := rand.Int()
	customer := model.NewCustomerModel(rand_id, name, gender, age, phone, email)
	this.customerController.Add(customer)
	fmt.Print("-- customer add end --")
}

func (this *CustomerView) delete() {
	fmt.Println("enter the user number to delete")
	var id int
	fmt.Scanln(&id)
	res := this.customerController.Delete(id)
	if res {
		fmt.Println("delete ok")
	} else {
		fmt.Println("delete fail")
	}
}

func (this *CustomerView) edit() {
	fmt.Println("-- edit customer --")
	var id int
	fmt.Println("enter the customer id to edit")
	fmt.Scanln(&id)

	fmt.Println("name:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("gender:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("phone:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("email:")
	email := ""
	fmt.Scanln(&email)

	fmt.Println("age:")
	age := 0
	fmt.Scanln(&age)

	this.customerController.Edit(id,name,gender,phone,age,email)
	fmt.Print("-- customer edit end --")
}

func (this *CustomerView) doExit(){
	this.loop = false
}

func main() {
	CustomerView := CustomerView{
		key:  "",
		loop: true,
	}
	CustomerView.customerController = controller.NewCustomerController()

	CustomerView.MainMenu()
}
