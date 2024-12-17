package objects

type Task struct {
	GID         string
	Title       string
	Description string
	StartDate   string
	DateDue     string
	Status      string
}

type Plan struct {
	GID         string
	Title       string
	Description string
	Handler     string
	Team        string
	Budget      string
}

type Project struct {
	GID         string
	Title       string
	Description string
	Handler     string
	Budget      string
	StartDate   string
	EndDate     string
}
