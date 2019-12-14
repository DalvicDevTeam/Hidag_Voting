package services

import (
	"github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/entity"
	"github.com/NatnaelBerhanu-1/Hackathon/e-voting-system/votingmachine"
)

// VMServiceImpl represents
type VMServiceImpl struct {
	vmRepo votingmachine.VMRepository
}

// NewVMServiceImpl creates
func NewVMServiceImpl(vmRepo votingmachine.VMRepository) *VMServiceImpl {
	return &VMServiceImpl{vmRepo: vmRepo}
}

// Authenticate performs user authentication
func (vmsi *VMServiceImpl) Authenticate(pcnum string) ([]entity.Party, []entity.Party, error) {
	regionalCands, nationalCands, err := vmsi.vmRepo.Authenticate(pcnum)

	if err != nil {
		return nil, nil, err
	}
	return regionalCands, nationalCands, nil
}

// Vote performs
func (vmsi *VMServiceImpl) Vote(rid int, nid int, vid string) error {
	err := vmsi.vmRepo.Vote(rid, nid, vid)
	if err != nil {
		return err
	}

	return nil
}
