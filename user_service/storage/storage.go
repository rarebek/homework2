package storage

import (
	"FreelancerMarketplace/user-service/pkg/logger"
	"FreelancerMarketplace/user-service/storage/postgres"
	"FreelancerMarketplace/user-service/storage/repo"
	"database/sql"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sql.DB
	userRepo repo.UserStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sql.DB, log logger.Logger) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db, log),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
