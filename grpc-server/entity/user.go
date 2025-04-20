package entity

type User struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}
