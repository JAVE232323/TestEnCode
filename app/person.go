package app

type Person struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
}
