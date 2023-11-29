package psql

import "github.com/jackc/pgx/v5/pgxpool"

type PsqlStorages struct {
	UserStorage *UserPsqlStorage
}

func New(connPull *pgxpool.Pool) PsqlStorages {
	return PsqlStorages{UserStorage: NewUserStorage(connPull)}
}
