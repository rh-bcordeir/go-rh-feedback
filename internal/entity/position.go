package entity

type Position struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}
