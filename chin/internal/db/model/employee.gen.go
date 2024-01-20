// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEmployee = "Employee"

// Employee mapped from table <Employee>
type Employee struct {
	EmployeeID int32     `gorm:"column:EmployeeId;primaryKey;autoIncrement:true" json:"EmployeeId"`
	LastName   string    `gorm:"column:LastName;not null" json:"LastName"`
	FirstName  string    `gorm:"column:FirstName;not null" json:"FirstName"`
	Title      string    `gorm:"column:Title" json:"Title"`
	ReportsTo  int32     `gorm:"column:ReportsTo" json:"ReportsTo"`
	BirthDate  time.Time `gorm:"column:BirthDate" json:"BirthDate"`
	HireDate   time.Time `gorm:"column:HireDate" json:"HireDate"`
	Address    string    `gorm:"column:Address" json:"Address"`
	City       string    `gorm:"column:City" json:"City"`
	State      string    `gorm:"column:State" json:"State"`
	Country    string    `gorm:"column:Country" json:"Country"`
	PostalCode string    `gorm:"column:PostalCode" json:"PostalCode"`
	Phone      string    `gorm:"column:Phone" json:"Phone"`
	Fax        string    `gorm:"column:Fax" json:"Fax"`
	Email      string    `gorm:"column:Email" json:"Email"`
}

// TableName Employee's table name
func (*Employee) TableName() string {
	return TableNameEmployee
}