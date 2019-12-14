package votingmachine

import "github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"

// VMService represents voting machine reated functinality imps
type VMService interface {
	Verify(pcid int) ([]entity.Party, []entity.Party, string, error) // int is voter id
	Authenticate(pcnum string) (entity.VoteMachine, error)
	Vote(rid int, nid int, vid string) error
}
