package entity

type UserRole string

type User struct {
	ID       uint
	Name     string
	LastName string
	Surname  string
	Phone    string
	Email    string
	Password string
	Role     UserRole
}
