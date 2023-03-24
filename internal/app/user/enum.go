package user

type Role int

const (
	UserRole Role = iota
	AdminRole
)

func (e Role) String() string {
	return [...]string{"user", "admin"}[e]
}
