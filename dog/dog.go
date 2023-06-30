package dog

type Dog struct {
	Id       int
	UserId   int
	Name     string
	isHungry bool
}

func New(userId int, name string) *Dog {
	return &Dog{UserId: userId, Name: name, isHungry: true}
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) GetId() int {
	return d.Id
}

func (d *Dog) Feed() {
	d.isHungry = false
}

func (d *Dog) IsHungry() bool {
	return d.isHungry
}
