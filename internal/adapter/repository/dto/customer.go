package dto

import "go-test-api/internal/model"

const (
	CustomerTableName = "customer"

	CustomerIDField      = "id"
	CustomerNameField    = "name"
	CustomerAgeField     = "age"
	CustomerAddressField = "address"
	CustomerSalaryField  = "salary"
	CustomerSectorField  = "sector"
)

type (
	CustomerDto struct {
		ID      int32  `gorm:"column:id"`
		Name    string `gorm:"column:name"`
		Age     int16  `gorm:"column:age"`
		Address string `gorm:"column:address"`
		Salary  int64  `gorm:"column:salary"`
		Sector  int32  `gorm:"column:sector"`
	}

	CustomersDto []CustomerDto
)

func (CustomerDto) TableName() string {
	return "customer"
}

func (c CustomerDto) FromModel(m *model.Customer) *CustomerDto {
	c.ID = m.ID
	c.Name = m.Name
	c.Age = m.Age
	c.Address = m.Address
	c.Salary = m.Salary
	c.Sector = m.Sector

	return &c
}

func (c CustomerDto) ToEntity() *model.Customer {
	doc := &model.Customer{
		ID:      c.ID,
		Name:    c.Name,
		Age:     c.Age,
		Address: c.Address,
		Salary:  c.Salary,
		Sector:  c.Sector,
	}

	return doc
}

func (cs CustomersDto) ToEntities() model.Customers {
	if cs == nil {
		return nil
	}

	var entities model.Customers
	for _, dto := range cs {
		entities = append(entities, dto.ToEntity())
	}

	return entities
}
