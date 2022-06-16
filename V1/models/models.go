package models

type User struct {
	Id       uint   `json:"id" gorm:"primary"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"-"`
	Email    string `gorm:"unique"`
}

type Product struct {
	Id   uint   `json:"id" gorm:"primary"`
	Name string `json:"name"`
	Stok int    `json:"stok"`
}

type Order struct {
	Id         uint   `json:"id" gorm:"primary"`
	Productid  string `json:"productid"`
	Userid     string `json:"userid"`
	Restoranid string `json:"restoranid"`
	ExpTime    string `json:"exptime"`
}

type Restoran struct {
	Id   uint   `json:"id" gorm:"primary"`
	Name string `json:"name"`
}

type Bill struct {
	Id         uint   `json:"id" gorm:"primary"`
	Restoranid string `json:"restoranid"`
	Userid     string `json:"userid"`
	Orderid    string `json:"orderid"`
}
