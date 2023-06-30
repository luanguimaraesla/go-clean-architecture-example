package dog

type ServiceRepository interface {
	GetDogsByUserId(int) []*Dog
}

type Service struct {
	repository ServiceRepository
}

func NewService(r ServiceRepository) *Service {
	return &Service{r}
}

func (s *Service) FeedUserDogs(userId int) {
	userDogs := s.repository.GetDogsByUserId(userId)

	for _, d := range userDogs {
		d.Feed()
	}
}

func (s *Service) GetUserDogsNames(userId int) []string {
	userDogs := s.repository.GetDogsByUserId(userId)

	names := make([]string, len(userDogs))
	for i, d := range userDogs {
		names[i] = d.GetName()
	}

	return names
}
