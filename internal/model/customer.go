package model

type (
	Customer struct {
		ID      int32  `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Age     int16  `json:"age,omitempty"`
		Address string `json:"address,omitempty"`
		Salary  int64  `json:"salary,omitempty"`
		Sector  int32  `json:"sector,omitempty"`
	}

	Customers []*Customer
)

type RequestGetCustomerByID struct {
	ID int32
}

type RequestUpdateCustomer struct {
	ID      int32
	Name    string
	Age     int16
	Address string
	Salary  int64
	Sector  int32
}
