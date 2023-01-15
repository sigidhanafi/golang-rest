package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	DeleteByID(ID int) error
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

func (s *service) DeleteByID(ID int) error {
	err := s.repository.DeleteByID(ID)

	return err
}
