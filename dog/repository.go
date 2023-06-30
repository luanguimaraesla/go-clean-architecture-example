package dog

var (
	dogCounter = 0
	dogs       = []*Dog{}
)

type DogRepository struct{}

func (r *DogRepository) CreateDog(d *Dog) {
	d.Id = dogCounter
	dogCounter++

	dogs = append(dogs, d)
}

func (r *DogRepository) ListDogs() []*Dog {
	return dogs
}

func (r *DogRepository) GetDogsByUserId(userId int) []*Dog {
	ds := []*Dog{}

	for _, d := range dogs {
		if d.UserId == userId {
			ds = append(ds, d)
		}
	}

	return ds
}
