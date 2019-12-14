package services

import (
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/votingmachine"
)

// VMServiceImpl represents
type VMServiceImpl struct {
	vmRepo votingmachine.VMRepository
}

// NewVMServiceImpl creates
func NewVMServiceImpl(vmRepo votingmachine.VMRepository) *VMServiceImpl {
	return &VMServiceImpl{vmRepo: vmRepo}
}

// Verify performs user authentication
func (vmsi *VMServiceImpl) Verify() ([]entity.Party, []entity.Party, string, error) {
	regionalCands, nationalCands, vid, err := vmsi.vmRepo.Verify()

	if err != nil {
		return nil, nil, "", err
	}
	return regionalCands, nationalCands, vid, nil
}

// Authenticate performs pc authentication
func (vmsi *VMServiceImpl) Authenticate(pcnum string) (entity.VoteMachine, error) {
	vm, err := vmsi.vmRepo.Authenticate(pcnum)

	if err != nil {
		return entity.VoteMachine{}, err
	}
	return vm, nil
}

// Vote performs
func (vmsi *VMServiceImpl) Vote(rid int, nid int, vid string) error {
	err := vmsi.vmRepo.Vote(rid, nid, vid)
	if err != nil {
		return err
	}

	return nil
}
