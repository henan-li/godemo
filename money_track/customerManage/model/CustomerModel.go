package model

import "fmt"

type CustomerModel struct {
	ID     int
	NAME   string
	GENDER string
	AGE    int
	PHONE  string
	EMAIL  string
}

func NewCustomerModel(id int, name string, gender string, age int, phone string, email string) CustomerModel {
	return CustomerModel{
		ID:     id,
		NAME:   name,
		GENDER: gender,
		AGE:    age,
		PHONE:  phone,
		EMAIL:  email,
	}
}

func (this CustomerModel) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.ID, this.NAME, this.GENDER, this.AGE, this.PHONE, this.EMAIL)
	return info
}
