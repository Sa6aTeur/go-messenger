package service

import "fmt"

type UserStorage interface {
	Create(s string)
}

type UserService struct {
	storage UserStorage
}

func NewUserService(r UserStorage) *UserService {
	return &UserService{storage: r}
}

func (s *UserService) Create(str string) {
	fmt.Println("GetOneById in Service", str)
	s.storage.Create(str)
}
