package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// AgentRepositoryImpl represents
type AgentRepositoryImpl struct {
	conn *sql.DB
}

// NewAgentRepositoryImpl creates
func NewAgentRepositoryImpl(conn *sql.DB) *AgentRepositoryImpl {
	return &AgentRepositoryImpl{conn: conn}
}

// Authenticate reps
func (ari *AgentRepositoryImpl) Authenticate(id string, password string) (entity.Agent, error) {
	fmt.Println(id, password)
	row := ari.conn.QueryRow("select * from agents where agent_id = $1 and password = $2", id, password)
	fmt.Println("row: ", row)
	agent := entity.Agent{}
	if row != nil {
		err := row.Scan(&agent.ID, &agent.PollNo, &agent.FirstName, &agent.LastName, &agent.Password, &agent.AgentID)
		if err != nil {
			return agent, err
		}
		return agent, nil
	} else {
		return agent, errors.New("agent not found")
	}
}

// Verify reps
func (ari *AgentRepositoryImpl) Verify(id string, pollnum string) ([]entity.VoteMachine, entity.Voter, error) {
	row := ari.conn.QueryRow("select * from voters where voter_id = $1 and poll_no = $2 and voted = 0", id, pollnum)
	voter := entity.Voter{}
	vms := []entity.VoteMachine{}
	if row != nil {
		row.Scan(&voter.ID, &voter.FirstName, &voter.LastName, &voter.SurName, &voter.IDNum, &voter.VoterID, &voter.PollNo, &voter.Disability, &voter.Voted)
		// rows, err := ari.conn.Query("select * from parties where id in (select party_id from party_poll where poll_id = (select id from polls where poll_no = $1))", voter.PollNo)
		rows, err := ari.conn.Query("select * from vote_machines where poll_no = $1 and status = 0", voter.PollNo)

		if err != nil {
			return vms, voter, errors.New("could not query the database")
		}
		for rows.Next() {
			vm := entity.VoteMachine{}
			err = rows.Scan(&vm.ID, &vm.PollNo, &vm.PcNo, &vm.Status)
			if err != nil {
				return vms, voter, err
			}

			vms = append(vms, vm)
		}

		return vms, voter, nil

	}
	return vms, voter, errors.New("voter not found")
}

// OpenPc reps
func (ari *AgentRepositoryImpl) OpenPc(voterid string, pcid int, lang string) error {

	_, err := ari.conn.Exec("insert into votem_voter(voter_id, votem_id, language) values($1, $2, $3)", voterid, pcid, lang)
	if err != nil {
		return errors.New("can't open pc")
	}

	_, err2 := ari.conn.Exec("update vote_machines set status = 1 where id = $1", pcid)
	if err2 != nil {
		return errors.New("can't open pc")
	}

	return nil
}
