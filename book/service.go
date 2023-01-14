package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()

	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	dataBook, err := s.repository.FindByID(ID)

	return dataBook, err
}
