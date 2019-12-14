package helper

import (
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// VMData represents vote machine data
type VMData struct {
	vm entity.VoteMachine
}

var vmd = entity.VoteMachine{}


// NewVMData create new vote machine 
func NewVMData(vm entity.VoteMachine) {
	vmd = vm
}

// GetVMData return the vote data
func GetVMData() *entity.VoteMachine {
	return &vmd
}