package auth

import "errors"

// Auther is a generic interface to implement password-based authentication and authorization
type Auther interface {
	Authenticate(user, pass string) (*User, error)
	Authorize(user *User, topic string, perm Permission) error
}

type Manager interface {
	AddUser(username, password string, role Role) error
	RemoveUser(username string) error
	Users() ([]*User, error)
	User(username string) (*User, error)
	ChangePassword(username, password string) error
	ChangeRole(username string, role Role) error
	AllowAccess(username string, topic string, read bool, write bool) error
	ResetAccess(username string, topic string) error
}

type User struct {
	Name   string
	Pass   string // hashed
	Role   Role
	Grants []Grant
}

type Grant struct {
	Topic string
	Read  bool
	Write bool
}

type Permission int

const (
	PermissionRead  = Permission(1)
	PermissionWrite = Permission(2)
)

type Role string

const (
	RoleAdmin = Role("admin")
	RoleUser  = Role("user")
	RoleNone  = Role("none")
)

var Everyone = &User{
	Name: "",
	Role: RoleNone,
}

var Roles = []Role{
	RoleAdmin,
	RoleUser,
	RoleNone,
}

func AllowedRole(role Role) bool {
	return role == RoleUser || role == RoleAdmin
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
)
