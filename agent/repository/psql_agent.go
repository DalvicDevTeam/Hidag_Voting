package repository

import (
	"database/sql"
	"errors"

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
	row := ari.conn.QueryRow("SELECT * FROM agents where agent_id = $1 and password =", id, password)

	agent := entity.Agent{}

	err := row.Scan(&agent.AgentID, &agent.FirstName, &agent.ID, &agent.LastName, &agent.Password, &agent.PollNo)
	if err != nil {
		return agent, err
	}

	if agent.Password != password {
		return agent, err
	}

	return entity.Agent{}, errors.New("")
}

// Verify reps
func (ari *AgentRepositoryImpl) Verify(id string) ([]entity.VoteMachine, entity.Voter, error) {
	return []entity.VoteMachine{}, entity.Voter{}, errors.New("")
}

// OpenPc reps
func (ari *AgentRepositoryImpl) OpenPc(voterid string, pcnum string) error {
	return errors.New("")
}
