package objects

type Member struct {
	ID             int
	UserID         int
	Date           string
	Name           string
	Dob            string
	Zone           string
	Village        string
	Parish         string
	Subcounty      string
	FormerReligion string
	Educ           string
	Occupation     string
	Where          string
	MaritalStatus  string
	Children       string
	Tel            string
	Email          string
	Father         string
	Mother         string
	HomeVillage    string
	HomeParish     string
	HomeSubCounty  string
	HomeDistrict   string
}

type Minister struct {
	ID             int
	UserID         int
	Date           string
	Name           string
	Ministry       string
	Role           string
	Dob            string
	Zone           string
	Village        string
	Parish         string
	Subcounty      string
	FormerReligion string
	Educ           string
	Occupation     string
	Where          string
	MaritalStatus  string
	Children       string
	Tel            string
	Email          string
	Father         string
	Mother         string
	HomeVillage    string
	HomeParish     string
	HomeSubCounty  string
	HomeDistrict   string
}

type Visitor struct {
	Tel          string
	Email        string
	HomeDistrict string
}

type Pastor struct {
	ID             int
	UserID         int
	Date           string
	Name           string
	Ministry       string
	Dob            string
	PastorSince    string
	Zone           string
	Village        string
	Parish         string
	Subcounty      string
	FormerReligion string
	Educ           string
	Occupation     string
	Where          string
	MaritalStatus  string
	Children       string
	Tel            string
	Email          string
	Father         string
	Mother         string
	HomeVillage    string
	HomeParish     string
	HomeSubCounty  string
	HomeDistrict   string
}

type Admin struct {
	GID        int
	Name       string
	Email      string
	ProfileImg string
}
