package entity

type Position struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
