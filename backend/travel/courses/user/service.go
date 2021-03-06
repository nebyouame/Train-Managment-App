package user

import "github.com/travel/courses/model"

type RoleService interface {
	Roles() ([]model.Role, []error)
	Role(id uint) (*model.Role, []error)
	RoleByName(name string) (*model.Role, []error)
	UpdateRole(role *model.Role) (*model.Role, []error)
	DeleteRole(id uint) (*model.Role, []error)
	StoreRole(role *model.Role) (*model.Role, []error)
}

type UserService interface {
	User(id uint) (*model.User, []error)
	UserByEmail(email string) (*model.User, []error)
	UpdateUser(user *model.User) (*model.User, []error)
	DeleteUser(id uint) (*model.User, []error)
	StoreUser(user *model.User) (*model.User, []error)
	UpdateUserAmount(user *model.User, Amount uint) *model.User
	EmailExists(email string) bool

}