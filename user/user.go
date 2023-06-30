package user

type DogService interface {
	FeedUserDogs(int)
	GetUserDogsNames(int) []string
}

type User struct {
	Id       int
	Username string
}

func New(name string) *User {
	return &User{Username: name}
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetName() string {
	return u.GetUsername()
}

func (u *User) FeedMyDogs(svc DogService) {
	svc.FeedUserDogs(u.Id)
}

func (u *User) GetUserDogsNames(svc DogService) []string {
	return svc.GetUserDogsNames(u.Id)
}
