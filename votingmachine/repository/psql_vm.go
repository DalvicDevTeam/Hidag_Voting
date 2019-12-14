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
	voter := entity.Votes{}
	_, err := vmri.conn.Exec("INSERT INTO votes(voter_id,regional_id,national_id) VALUES($1,$2,$3)", voter.VoterID, voter.RegionalID, voter.NationalID)

	if err != nil {

		return errors.New("")
	}

	_, err = vmri.conn.Exec("UPDATE party_poll SET regional_vote=regional_vote + 1 WHERE party_id = $1 ", rid)
	_, err = vmri.conn.Exec("UPDATE party_poll SET national_vote=national_vote + 1 WHERE party_id = $1 ", nid)
	_, err = vmri.conn.Exec("DELETE FROM votes WHERE voter_id = $1;", vid)
	return errors.New("")
}
