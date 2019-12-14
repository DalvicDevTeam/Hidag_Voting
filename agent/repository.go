package agent

import (
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// AgentRepository rep
type AgentRepository interface {
	Authenticate(id string, password string) (entity.Agent, error)
	Verify(id string) ([]entity.VoteMachine, entity.Voter, error)
	OpenPc(voterid string, pcnum string, lang string) error
}
