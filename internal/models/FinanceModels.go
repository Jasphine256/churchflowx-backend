package models

type Payment struct {
	ID     int    `gorm:"primaryKey"`
	GID    int    `gorm:"not null"`
	Name   string `gorm:"size:255;not null"`
	Reason string `gorm:"size:255;not null"`
	Amount string `gorm:"size:50;not null"`
	Date   string `gorm:"size:255;not null"`
}

type Fund struct {
	ID     int    `gorm:"primaryKey"`
	GID    int    `gorm:"not null"`
	Name   string `gorm:"size:255;not null"`
	Reason string `gorm:"size:255;not null"`
	Amount string `gorm:"size:50;null"`
	Date   string `gorm:"size:255;not null"`
}
