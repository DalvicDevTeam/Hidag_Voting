package votingmachine

import "github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"

// VMRepository represents voting machine reated functinality imps
type VMRepository interface {
	Verify(pcid int) ([]entity.Party, []entity.Party, string, error) // int is voter id
	Authenticate(pcnum string) ([]entity.Party, []entity.Party, entity.VoteMachine, error)
	Vote(rid int, nid int, vid int) error
}