package main

type user struct {
	Username    string 	`bson:"Username"`
	Password	string  `bson:"Password"`	
	UserId		string  `bson:"UserId"`
}

type item struct {
	Name		string  `bson:"Name"`
	Description	string	`bson:"Description"`
	Id			string	`bson:"Id"`
	Quantity	int		`bson:"Quantity"`
	Price		int		`bson:"Price"`
}

type orderDetails struct {
	OrderId     string 		`bson:"OrderId"`
	Items		item    	`bson:"Items"`
	UserId		string		`bson:"UserId"`
}

type cart struct {
	UserId 	string	`bson:"UserId"`
	CartId	string	`bson:"CartId"`
	Status	string	`bson:"Status"`
	Orders 	[]orderDetails	`bson:"Orders"`
}

type paymentDetails struct {
	PaymentId   	string 	
	OrderId			string    	
	UserId			string
	FullName		string
	Phone			string
	PaymentType		string
	TotalPrice		int
}