package models

type Task struct {
	ID          int    `gorm:"primaryKey"`
	GID         int    `gorm:"not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:1000;not null"`
	StartDate   string `gorm:"size:25;not null"`
	DateDue     string `gorm:"size:25;not null"`
	Status      string `gorm:"size:25;not null"`
}

type Plan struct {
	ID          int    `gorm:"primaryKey"`
	GID         int    `gorm:"not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:1000;not null"`
	Handler     string `gorm:"size:25;not null"`
	Team        string `gorm:"size:25;not null"`
	Budget      string `gorm:"size:50;not null"`
}

type Project struct {
	ID          int    `gorm:"primaryKey"`
	GID         int    `gorm:"not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:1000;not null"`
	Handler     string `gorm:"size:25;not null"`
	Budget      string `gorm:"size:50;not null"`
	StartDate   string `gorm:"size:25;not null"`
	EndDate     string `gorm:"size:25"`
}
