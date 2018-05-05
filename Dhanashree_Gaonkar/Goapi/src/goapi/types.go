/*
	Starbucks API in Go (Version 3)
	Uses MongoDB and Redis 
	(For use with Kong API Key)
*/
	
package main

type test struct {
		    
	UserId 	string
	Id             	int 	
	Count   int    	
	ModelNumber 	string	    
	SerialNumber 	string	
}

type user struct {
	Email       string 
	Password	string    	
	UserId		string
}

type item struct {
	Name		string
	Description	string
	Id			string
	Quantity	int
	Price		int
}

type orderDetails struct {
	OrderId     string 	
	Items		item    	
	UserId		string
}

type cart struct {
	UserId 	string
	CartId	string
	Status	string
	orders 	[]orderDetails
}


type paymentDetails struct {
	PaymentId   	string 	
	CartId			string    	
	UserId			string
	PaymentType		string
	TotalPrice		int
}
