package ports

type UserService interface {
	Login(username string, password string) (string, error)
	Signup(email string, password string) (string, error)
}
