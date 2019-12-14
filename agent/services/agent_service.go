package services

import (
	"github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/agent"
	"github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/entity"
)

// AgentServiceImpl represents
type AgentServiceImpl struct {
	agentRepo agent.AgentRepository
}

// NewAgentServiceImpl creates
func NewAgentServiceImpl(agentRepo agent.AgentRepository) *AgentServiceImpl {
	return &AgentServiceImpl{agentRepo: agentRepo}
}

// Authenticate reps
func (asi *AgentServiceImpl) Authenticate(id string, password string) (entity.Agent, error) {
	agent, err := asi.agentRepo.Authenticate(id, password)

	if err != nil {
		return entity.Agent{}, err
	}

	return agent, nil
}

// Verify reps
func (asi *AgentServiceImpl) Verify(id string) ([]entity.VoteMachine, entity.Voter, error) {
	vm, voter, err := asi.agentRepo.Verify(id)

	if err != nil {
		return nil, entity.Voter{}, err
	}

	return vm, voter, nil
}

// OpenPc reps
func (asi *AgentServiceImpl) OpenPc(voterid string, pcnum string, lang string) error {
	err := asi.agentRepo.OpenPc(voterid, pcnum, lang)

	if err != nil {
		return err
	}

	return nil
}