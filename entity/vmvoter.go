package entity

// VMVoter represents temporary mapping between voice machine and voter
type VMVoter struct {
	ID            int
	VoterID       string
	VoteMachineID int
	Language      string
}
