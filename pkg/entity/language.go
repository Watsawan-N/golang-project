package entity

type Language struct {
	DefaultField
	Code    *string
	Name    *string
	SAPCode *string
}

func (Language) TableName() string {
	return "system_languages"
}
