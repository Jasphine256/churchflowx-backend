package models

type Payment struct {
	ID     int    `gorm:"primaryKey"`
	UserId int    `gorm:"not null"`
	Name   string `gorm:"size:255;not null"`
	Reason string `gorm:"size:255;not null"`
	Amount uint16 `gorm:"not null"`
	Date   string `gorm:"size:255;not null"`
}

type Fund struct {
	ID     int    `gorm:"primaryKey"`
	UserId int    `gorm:"not null"`
	Name   string `gorm:"size:255;not null"`
	Reason string `gorm:"size:255;not null"`
	Amount uint16 `gorm:"not null"`
	Date   string `gorm:"size:255;not null"`
}
