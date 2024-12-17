package models

type Member struct {
	ID             int    `gorm:"primaryKey"`
	GID            string `gorm:"not null"`
	Date           string `gorm:"size:25;not null"`
	Name           string `gorm:"size:100;not null"`
	Dob            string `gorm:"size:25;not null"`
	Zone           string `gorm:"size:100;not null"`
	Village        string `gorm:"size:100;not null"`
	Parish         string `gorm:"size:100;not null"`
	Subcounty      string `gorm:"size:100;not null"`
	FormerReligion string `gorm:"size:50;not null"`
	Educ           string `gorm:"size:100;not null"`
	Occupation     string `gorm:"size:50;not null"`
	Where          string `gorm:"size:100;not null"`
	MaritalStatus  string `gorm:"size:25;not null"`
	Children       string `gorm:"ize:10;not null"`
	Tel            string `gorm:"size:255;not null"`
	Email          string `gorm:"size:255;not null"`
	Father         string `gorm:"size:100;not null"`
	Mother         string `gorm:"size:100;not null"`
	HomeVillage    string `gorm:"size:100;not null"`
	HomeParish     string `gorm:"size:100;not null"`
	HomeSubCounty  string `gorm:"size:100;not null"`
	HomeDistrict   string `gorm:"size:100;not null"`
}

type Minister struct {
	ID             int    `gorm:"primaryKey"`
	GID            string `gorm:"not null"`
	Date           string `gorm:"size:25;not null"`
	Name           string `gorm:"size:100;not null"`
	Ministry       string `gorm:"size:100;not null"`
	Role           string `gorm:"size:25;not null"`
	Dob            string `gorm:"size:25;not null"`
	Zone           string `gorm:"size:100;not null"`
	Village        string `gorm:"size:100;not null"`
	Parish         string `gorm:"size:100;not null"`
	Subcounty      string `gorm:"size:100;not null"`
	FormerReligion string `gorm:"size:50;not null"`
	Educ           string `gorm:"size:100;not null"`
	Occupation     string `gorm:"size:50;not null"`
	Where          string `gorm:"size:100;not null"`
	MaritalStatus  string `gorm:"size:25;not null"`
	Children       string `gorm:"size:10;not null"`
	Tel            string `gorm:"size:255;not null"`
	Email          string `gorm:"size:255;not null"`
	Father         string `gorm:"size:100;not null"`
	Mother         string `gorm:"size:100;not null"`
	HomeVillage    string `gorm:"size:100;not null"`
	HomeParish     string `gorm:"size:100;not null"`
	HomeSubCounty  string `gorm:"size:100;not null"`
	HomeDistrict   string `gorm:"size:100;not null"`
}

type Visitor struct {
	ID           int    `gorm:"primaryKey"`
	GID          string `gorm:"not null"`
	Tel          string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;not null"`
	HomeDistrict string `gorm:"size:25;not null"`
}

type Pastor struct {
	ID             int    `gorm:"primaryKey"`
	GID            string `gorm:"not null"`
	Date           string `gorm:"size:25;not null"`
	Name           string `gorm:"size:100;not null"`
	Ministry       string `gorm:"size:100;not null"`
	Dob            string `gorm:"size:25;not null"`
	PastorSince    string `gorm:"size:25;not null"`
	Zone           string `gorm:"size:100;not null"`
	Village        string `gorm:"size:100;not null"`
	Parish         string `gorm:"size:100;not null"`
	Subcounty      string `gorm:"size:100;not null"`
	FormerReligion string `gorm:"size:50;not null"`
	Educ           string `gorm:"size:100;not null"`
	Occupation     string `gorm:"size:50;not null"`
	Where          string `gorm:"size:100;not null"`
	MaritalStatus  string `gorm:"size:25;not null"`
	Children       string `gorm:"size:10;not null"`
	Tel            string `gorm:"size:255;not null"`
	Email          string `gorm:"size:255;not null"`
	Father         string `gorm:"size:100;not null"`
	Mother         string `gorm:"size:100;not null"`
	HomeVillage    string `gorm:"size:100;not null"`
	HomeParish     string `gorm:"size:100;not null"`
	HomeSubCounty  string `gorm:"size:100;not null"`
	HomeDistrict   string `gorm:"size:100;not null"`
}

type Admin struct {
	GID        string     `gorm:"primaryKey"`
	Name       string     `gorm:"size:100;not null"`
	Email      string     `gorm:"size:100;not null"`
	ProfileImg string     `gorm:"size:255;not null"`
	Members    []Member   `gorm:"foreignKey:GID"`
	Ministers  []Minister `gorm:"foreignKey:GID"`
	Visitors   []Visitor  `gorm:"foreignKey:GID"`
	Pastors    []Pastor   `gorm:"foreignKey:GID"`
	Tasks      []Task     `gorm:"foreignKey:GID"`
	Plans      []Plan     `gorm:"foreignKey:GID"`
	Projects   []Project  `gorm:"foreignKey:GID"`
	Funds      []Fund     `gorm:"foreignKey:GID"`
	Payments   []Payment  `gorm:"foreignKey:GID"`
}
