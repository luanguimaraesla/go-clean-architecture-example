package user

var (
	userCounter = 0
	users       = []*User{}
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(d *User) {
	d.Id = userCounter
	userCounter++

	users = append(users, d)
}

func (r *UserRepository) ListUsers() []*User {
	return users
}
