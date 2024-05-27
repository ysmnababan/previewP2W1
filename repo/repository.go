package repo

import (
	"database/sql"
	"errors"
	"log"
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
	AddNewBranch(b model.Branch) (model.Branch, error)
	UpdateBranch(id int, b model.Branch) error
	DeleteBranch(id int) error
}

type MysqlRepo struct {
	DB *sql.DB
}

func (r *MysqlRepo) isIDExist(id int) (bool, error) {
	isExist := true

	err := r.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM branches WHERE branch_id = ?)", id).Scan(&isExist)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *MysqlRepo) GetAllBranches() ([]model.Branch, error) {
	var branches []model.Branch

	rows, err := r.DB.Query("SELECT branch_id, name, location FROM branches")
	if err != nil {
		log.Println("error query")
		return nil, ErrQuery
	}
	defer rows.Close()

	for rows.Next() {
		var b model.Branch
		err = rows.Scan(&b.Branch_Id, &b.Name, &b.Location)
		if err != nil {
			log.Println("error scan row")
			return nil, ErrScan
		}

		branches = append(branches, b)
	}

	return branches, nil
}

func (r *MysqlRepo) GetBranchByID(id int) (model.Branch, error) {
	var b model.Branch

	isExist, _ := r.isIDExist(id)
	if !isExist {
		log.Println("branch is not exist: ", id)
		return model.Branch{}, ErrNoRows
	}

	err := r.DB.QueryRow("SELECT name, location FROM branches WHERE branch_id = ?", id).Scan(&b.Name, &b.Location)
	if err != nil {
		log.Println("error scan row")
		return model.Branch{}, ErrScan
	}

	b.Branch_Id = id
	return b, nil
}

func (r *MysqlRepo) AddNewBranch(b model.Branch) (model.Branch, error) {
	result, err := r.DB.Exec("INSERT INTO branches (name, location) VALUES (?,?)", b.Name, b.Location)
	if err != nil {
		log.Println("error query")
		return model.Branch{}, ErrQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("error getting last id")
		return model.Branch{}, ErrLastInsertId
	}

	b.Branch_Id = int(lastID)
	return b, nil
}

func (r *MysqlRepo) UpdateBranch(id int, b model.Branch) error {
	isExist, _ := r.isIDExist(id)
	if !isExist {
		log.Println("branch is not exist: ", id)
		return ErrNoRows
	}

	result, err := r.DB.Exec("UPDATE branches SET name=?, location=? WHERE branch_id = ?", b.Name, b.Location, id)
	if err != nil {
		log.Println("error query")
		return ErrQuery
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		log.Println("error getting rows affected")
		return ErrRowsAffected
	}

	if affectedRow == 0 {
		log.Println("no updated rows")
		return ErrNoAffectedRow
	}

	return nil
}

func (r *MysqlRepo) DeleteBranch(id int) error {
	isExist, _ := r.isIDExist(id)
	if !isExist {
		log.Println("branch is not exist: ", id)
		return ErrNoRows
	}

	result, err := r.DB.Exec("DELETE FROM branches WHERE branch_id= ?", id)
	if err != nil {
		log.Println("error query")
		return ErrQuery
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		log.Println("error getting rows affected")
		return ErrRowsAffected
	}

	if affectedRow == 0 {
		log.Println("no updated rows")
		return ErrNoAffectedRow
	}
	return nil
}
