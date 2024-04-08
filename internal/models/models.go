package models

type Car struct {
	Nomer   string `json:"regNum" db:"regnum"`
	Mark    string `json:"mark" db:"mark"`
	Model   string `json:"model" db:"model"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}

type CarUpdate struct {
	Nomer   *string `json:"regNum" db:"regnum"`
	Mark    *string `json:"mark" db:"mark"`
	Model   *string `json:"model" db:"model"`
	Name    *string `json:"name" db:"name"`
	Surname *string `json:"surname" db:"surname"`
}
