package model

type User struct {
	Id         *uint   `json:"id"`
	UserId     *string `json:"userId"`
	EmployeeId *string `json:"employeeId"`
	Remark     *string `json:"remark"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	CreatedBy  *uint   `json:"createdBy"`
	UpdatedBy  *uint   `json:"updatedBy"`
	PersonId   *string `json:"personId"`
}
