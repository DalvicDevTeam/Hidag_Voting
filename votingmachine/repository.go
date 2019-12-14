package votingmachine

import "github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/entity"

// VMRepository represents voting machine reated functinality imps
type VMRepository interface {
	Authenticate(pcnum string) ([]entity.Party, []entity.Party, error)
	Vote(rid int, nid int, vid string) error
}