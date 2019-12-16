package helper

import (
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"
)

// VMData represents vote machine data
type VMData struct {
	vm entity.VoteMachine
}

//AgentData struct
type AgentData struct {
	agent entity.Agent
}

var Vmd = []entity.VoteMachine{}
var voter = entity.Voter{}
var Agents = []entity.Agent{}

// NewVMData create new vote machine
func NewVMData(vm entity.VoteMachine) {
	Vmd = append(Vmd, vm)
}

// NewVoter create new vote machine
func NewVoter(voterd entity.Voter) {
	voter = voterd
}

// GetVMData return the vote data
func GetVMData() entity.VoteMachine {
	return Vmd[0]
}

// GetVoter return the vote data
func GetVoter() *entity.Voter {
	return &voter
}

// NewAgentData create new agent data
func NewAgentData(vm entity.Agent) {
	Agents = append(Agents, vm)
}

// GetAgentData return
func GetAgentData() entity.Agent {
	return Agents[0]
}
