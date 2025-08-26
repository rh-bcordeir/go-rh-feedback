package entity

type Stage struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
}
