package handler

import (
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/votingmachine"
)

//VMHandler handles
type VMHandler struct {
	tmpl  *template.Template
	vmSrv votingmachine.VMService
}

// NewVMHandler creates new Voter handler struct
func NewVMHandler(T *template.Template, vmSrv votingmachine.VMService) *VMHandler {
	return &VMHandler{tmpl: T, vmSrv: vmSrv}
}

// VoterIndex handles request on /voter
func (ah *VMHandler) VoterIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "voter.vote.layout", nil)
	}
}

// VoterLogin handles request on /voter/login
func (ah *VMHandler) VoterLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "voter.login.layout", nil)
	}
}

// VoterVote handles request on /voter/vote
func (ah *VMHandler) VoterVote(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "voter.vote.layout", nil)
	}
}
