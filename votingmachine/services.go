package votingmachine

import "github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"

// VMService represents voting machine reated functinality imps
type VMService interface {
	Authenticate(pcnum string) ([]entity.Party, []entity.Party, error)
	Vote(rid int, nid int) error
}
