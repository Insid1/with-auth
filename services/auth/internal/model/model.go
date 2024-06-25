package model

type Login struct {
	Email    string
	ID       string
	Password string
}

type Register struct {
	Email    string
	Password string
}

type Check struct {
	AccessToken  string
	RefreshToken string
}
