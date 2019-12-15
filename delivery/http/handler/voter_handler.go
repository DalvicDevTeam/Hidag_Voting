package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/helper"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/votingmachine"
)

//VMHandler handles
type VMHandler struct {
	tmpl  *template.Template
	vmSrv votingmachine.VMService
}

// PData struct
type PData struct {
	RegCand []entity.Party
	NatCand []entity.Party
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
	} else {
		vmnum := r.FormValue("vmnum")

		fmt.Println("pcnum: ", vmnum)

		regionalCands, nationalCands, vm, err := ah.vmSrv.Authenticate(vmnum)

		data := &PData{RegCand: regionalCands, NatCand: nationalCands}

		if err != nil {
			panic(err)
		}

		helper.NewVMData(vm)
		ah.tmpl.ExecuteTemplate(w, "voter.vote.layout", data)
	}
}

// VoterVote handles request on /voter/vote
func (ah *VMHandler) VoterVote(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "voter.vote.layout", nil)
	} else {
		natid, err := strconv.Atoi(r.PostFormValue("nid"))
		regid, err1 := strconv.Atoi(r.PostFormValue("rid"))
		if err != nil && err1 != nil {

		}
		vid := helper.GetVMData().ID

		err2 := ah.vmSrv.Vote(regid, natid, vid)

		if err2 != nil {
			panic(err2)
			w.Write([]byte("error"))
		} else {
			w.Write([]byte("success"))
		}

	}
}
