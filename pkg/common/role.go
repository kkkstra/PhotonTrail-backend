package common

type Role int

const (
	Admin Role = iota
	User
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case User:
		return "users"
	default:
		return "unknown"
	}
}
