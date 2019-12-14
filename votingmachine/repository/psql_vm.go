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

// Verify performs user verification
func (vmri *VMRepositoryImpl) Verify(pcid int) ([]entity.Party, []entity.Party, string, error) {
	row := vmri.conn.QueryRow("select * from votem_voter where votem_id = $1", pcid)
	vm := entity.VMVoter{}
	if row != nil {
		err := row.Scan(&vm.ID, &vm.VoterID, &vm.VoteMachineID, &vm.Language)

		pollidrow := vmri.conn.QueryRow("select * from voters where voter_id = $1", vm.VoterID)

		if err != nil {
			return nil, nil, "", errors.New(" votem voter Row scannng failed")
		}

		regrows, err1 := vmri.conn.Query("select * from parties where id in (select party_id from party_poll where poll_id = (select id from polls where poll_no = $1)) and is_regional != 0", vm.PollNo)
		natrows, err2 := vmri.conn.Query("select * from parties where id in (select party_id from party_poll where poll_id = (select id from polls where poll_no = $1)) and is_regional != 0", vm.PollNo)

		regionalCands := []entity.Party{}
		nationalCands := []entity.Party{}

		if err1 != nil || err2 != nil {
			for regrows.Next() {
				rc := entity.Party{}
				err := regrows.Scan(&rc.ID, &rc.ImageURL, &rc.Representative, &rc.Name)

				if err != nil {
					return nil, nil, "", errors.New(" regrows scanning failed")
				}
				regionalCands = append(regionalCands, rc)
			}
			for natrows.Next() {
				nc := entity.Party{}
				err := regrows.Scan(&nc.ID, &nc.ImageURL, &nc.Representative, &nc.Name)

				if err != nil {
					return nil, nil, "", errors.New(" regrows scanning failed")
				}
				nationalCands = append(nationalCands, nc)
			}

			return regionalCands, nationalCands, vm.VoterID, nil
		}

	}
	return nil, nil, "", errors.New("Verification failed")
}

// Authenticate performs user authentication
func (vmri *VMRepositoryImpl) Authenticate(pcnum string) (entity.VoteMachine, error) {
	row := vmri.conn.QueryRow("select * from vote_machines where pc_num = $1", pcnum)
	vm := entity.VoteMachine{}
	if row != nil {
		err := row.Scan(&vm.ID, &vm.PollNo, &vm.PcNo, &vm.Status)
		if err != nil {
			return vm, errors.New("Authentication failed")
		}
	}

	return vm, nil
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
