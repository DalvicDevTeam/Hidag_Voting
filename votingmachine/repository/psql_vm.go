package repository

import (
	"database/sql"
	"errors"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// VMRepositoryImpl represents
type VMRepositoryImpl struct {
	conn *sql.DB
}

// NewVMRepositoryImpl creates
func NewVMRepositoryImpl(conn *sql.DB) *VMRepositoryImpl {
	return &VMRepositoryImpl{conn: conn}
}

// Authenticate performs user authentication
func (vmri *VMRepositoryImpl) Authenticate(pcnum string) ([]entity.Party, []entity.Party, error) {
	return []entity.Party{}, []entity.Party{}, errors.New("")
}

// Vote performs
func (vmri *VMRepositoryImpl) Vote(rid int, nid int, vid string) error {
	return errors.New("")
}
