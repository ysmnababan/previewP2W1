package repo

import (
	"database/sql"
	"errors"
	"pagi/model"
)

var (
	ErrNoRows        = errors.New("no rows in result set")
	ErrQuery         = errors.New("query execution failed")
	ErrScan          = errors.New("row scanning failed")
	ErrInvalidId     = errors.New("invalid id")
	ErrUserExists    = errors.New("user already exist")
	ErrRowsAffected  = errors.New("unable to get affected row")
	ErrNoAffectedRow = errors.New("rows affected is 0")
	ErrLastInsertId  = errors.New("unable to get last insert id")
)

type GameStoreRepo interface {
	GetAllBranches() ([]model.Branch, error)
	GetBranchByID(id int) (model.Branch, error)
	AddNewBranch() (model.Branch, error)
	UpdateBranch(id int, b model.Branch) (model.Branch, error)
	DeleteBranch(id int) error
}

type MysqlRepo struct {
	DB *sql.DB
}

func (r *MysqlRepo) GetAllBranches() ([]model.Branch, error) {
	return nil, nil
}

func (r *MysqlRepo) GetBranchByID(id int) (model.Branch, error) {

	return model.Branch{}, nil
}

func (r *MysqlRepo) AddNewBranch() (model.Branch, error) {

	return model.Branch{}, nil
}

func (r *MysqlRepo) UpdateBranch(id int, b model.Branch) (model.Branch, error) {
	return model.Branch{}, nil
}

func (r *MysqlRepo) DeleteBranch(id int) error {
	return nil
}
