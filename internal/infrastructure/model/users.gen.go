// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	Role     string `gorm:"column:role" json:"role"`
	Username string `gorm:"column:username" json:"username"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
