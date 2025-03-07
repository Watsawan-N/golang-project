package entity

type EventLang struct {
	DefaultField
	EventId    *uint
	LanguageId *uint
	Language   *Language `gorm:"foreignKey:LanguageId"`
	Name       *string
}

func (EventLang) TableName() string {
	return "master_event_langs"
}
