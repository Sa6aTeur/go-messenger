package psql

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPsqlStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(cp *pgxpool.Pool) *UserPsqlStorage {
	return &UserPsqlStorage{db: cp}
}

func (s *UserPsqlStorage) Create(str string) {
	fmt.Println("GetOneById in storage", str)
}
