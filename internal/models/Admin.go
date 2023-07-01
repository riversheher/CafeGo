package models

type Admin struct {
	User
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateAdminTable() {

}
