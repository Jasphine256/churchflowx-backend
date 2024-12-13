package objects

type Task struct {
	Title       string
	Description string
	StartDate   string
	DateDue     string
	Status      string
}

type Plan struct {
	Title       string
	Description string
	Handler     string
	Team        string
	Budget      string
}

type Project struct {
	Title       string
	Description string
	Handler     string
	Budget      string
	StartDate   string
	EndDate     string
}
