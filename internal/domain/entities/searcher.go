package entities

type Search struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Language         string `json:"language"`
	Status           Status `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
