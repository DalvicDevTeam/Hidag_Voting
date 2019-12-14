package votingmachine

import "github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/entity"

// VMService represents voting machine reated functinality imps
type VMService interface {
	Authenticate(pcnum string) ([]entity.Party, []entity.Party, error)
	Vote(rid int, nid int) error
}
