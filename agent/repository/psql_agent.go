package repository

import (
	"database/sql"
	"errors"

	"github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/entity"
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
