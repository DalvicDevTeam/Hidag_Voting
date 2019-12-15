package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/entity"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/helper"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent"
)

//AgentHandler handles
type AgentHandler struct {
	tmpl     *template.Template
	agentSrv agent.AgentService
}

// Data struct
type Data struct {
	Vms   []entity.VoteMachine
	Voter entity.Voter
}

// NewAgentHandler creates new agent handler struct
func NewAgentHandler(T *template.Template, catserv agent.AgentService) *AgentHandler {
	return &AgentHandler{tmpl: T, agentSrv: catserv}
}

// AgentIndex handles request on /agent
func (ah *AgentHandler) AgentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.verify.layout", nil)
	}
}

// AgentLogin handles request on /agent/login
func (ah *AgentHandler) AgentLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.login.layout", nil)
	} else {
		id := r.FormValue("id")
		pass := r.FormValue("password")
		fmt.Println(id, pass)
		agent, err := ah.agentSrv.Authenticate(id, pass)
		helper.NewAgentData(agent)
		if err != nil {
			panic(err)
		}
		ah.tmpl.ExecuteTemplate(w, "agent.verify.layout", agent)

	}
}

// AgentVerifyVoter handles request on /agent/verifyvoter
func (ah *AgentHandler) AgentVerifyVoter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.verified.layout", nil)
		fmt.Println(helper.GetAgentData())
	} else {
		uuid := r.FormValue("uuid")
		pollnum := r.FormValue("pollnum")

		fmt.Println("uuid: ", uuid)
		fmt.Println("pollnum: ", pollnum)

		vms, voter, err := ah.agentSrv.Verify(uuid, pollnum)
		data := &Data{Vms: vms, Voter: voter}
		helper.NewVoter(voter)
		if err != nil {
			panic(err)
		}

		ah.tmpl.ExecuteTemplate(w, "agent.verified.layout", data)
		// http.Redirect(w, r, "/agent/verified", http.StatusSeeOther)
	}
}

// Openpc handler request on agent/openpc
func (ah *AgentHandler) Openpc(w http.ResponseWriter, r *http.Request) {
	voterid := r.PostFormValue("voterid")
	pcn := r.PostFormValue("pcnum")
	// w.Write([]byte("failed"))
	// fmt.Println(voterid, pcn)
	pcnum, err := strconv.Atoi(pcn)
	if err != nil {
		panic(err)
	}
	err2 := ah.agentSrv.OpenPc(voterid, pcnum, "")
	if err2 != nil {
		panic(err)
	}
}
