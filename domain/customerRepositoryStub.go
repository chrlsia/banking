package domain

type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer,error){
	return s.customers,nil
}

// Create all dummy customers
func NewCustomerRepositoryStub() CustomerRepositoryStub{
	customers:=[]Customer{
		{Id: "1001",Name: "Asish",City: "New Delhi",Zipcode: "110011",DateOfBirth: "2001-01-01", Status: "1"},
		{Id: "1002",Name: "Rob",City: "New Delhi",Zipcode: "110011",DateOfBirth: "2001-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}