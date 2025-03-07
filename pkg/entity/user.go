package entity

type User struct {
	DefaultField
	PersonId   *string `gorm:"type:nvarchar(20)"`
	UserId     *string `gorm:"type:nvarchar(20)"`
	EmployeeId *string `gorm:"type:nvarchar(20)"`
	Remark     *string `gorm:"type:nvarchar(500)"`
	FirstName  *string
	LastName   *string
}

func (User) TableName() string {
	return "system_users"
}
