package user

type Role int

const (
	RoleUser Role = iota
	RoleAdmin
)

func (e Role) String() string {
	return [...]string{"user", "admin"}[e]
}
