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
