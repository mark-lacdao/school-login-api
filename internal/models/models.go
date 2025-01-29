package models

// Student represents a student in the school
type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Grade     string `json:"grade"`
}

// Teacher represents a teacher in the school
type Teacher struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Subject   string `json:"subject"`
}
