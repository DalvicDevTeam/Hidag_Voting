package agent

import (
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// AgentService rep
type AgentService interface {
	Authenticate(id string, password string) (entity.Agent, error)
	Verify(id string) ([]entity.VoteMachine, entity.Voter, error)
	OpenPc(voterid string, pcid int, lang string) error
}
