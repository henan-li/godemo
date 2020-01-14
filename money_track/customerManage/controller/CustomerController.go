package controller

import (
	"../model"
)

type CustomerController struct {

	customers   []model.CustomerModel
	customerNum int
}

func NewCustomerController() *CustomerController {
	this := &CustomerController{}
	this.customerNum = 1

	customer := model.NewCustomerModel(1, "a", "male", 20, "321321", "test@test")
	this.customers = append(this.customers, customer)
	return this
}


func (this *CustomerController) List() []model.CustomerModel{

	return this.customers
}

func (this *CustomerController) Add(customer model.CustomerModel) bool{
	this.customerNum++
	customer.ID = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

func (this *CustomerController) findById(id int) int{

	num := -1
	for i:=0;i<len(this.customers) ;i++  {
		if this.customers[i].ID == id {
			num = i
		}
	}
	return num
}

func (this *CustomerController) Delete(id int) bool{

	index := this.findById(id)
	if index == -1{
		return false
	}
	for i:=0;i<len(this.customers) ;i++  {
		if this.customers[i].ID == id{
			this.customers = append(this.customers[:index],this.customers[index+1:]...)
		}
	}
	return true
}

func (this *CustomerController) Edit(id int,name string,gender string,phone string,age int,email string) bool{

	index := this.findById(id)
	if index == -1{
		return false
	}
	for i:=0;i<len(this.customers) ;i++  {
		if this.customers[i].ID == id{
			this.customers[i].NAME = name
			this.customers[i].GENDER = gender
			this.customers[i].PHONE = phone
			this.customers[i].AGE = age
			this.customers[i].EMAIL = email
		}
	}
	return true
}