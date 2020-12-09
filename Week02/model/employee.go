package model

type Employee struct {
	BaseModel
	Name      string `gorm: "column:name" json:"name" binding:"required"`
	Address   string `gorm: "column:address" json:"address"`
	Telephone string `gorm: "column:telephone" json:"telephone"`
	Email     string `gorm: "column:email" json:"email"`
}
