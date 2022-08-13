package entity

type Task struct {
	ID            uint   `json:"id"`
	CategoryId    int    `json:"category_id"`
	DefinedBy     int    `json:"defined_by"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	DefinedPerson User   `json:"defined_person"`
	AssignTo      []User `gorm:"many2many:user_languages;"`
}
