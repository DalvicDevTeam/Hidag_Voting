package handler

import (
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent"
)

//AgentHandler handles
type AgentHandler struct {
	tmpl     *template.Template
	agentSrv agent.AgentService
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
		id := r.FormValue("email")
		pass := r.FormValue("password")
		_, err := ah.agentSrv.Authenticate(id, pass)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/Agent", http.StatusSeeOther)

	}
}

// AgentVerified handles request on /agent/verified
func (ah *AgentHandler) AgentVerified(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.verified.layout", nil)
	} else {
		uuid := r.FormValue("uuid")
		_, _, err := ah.agentSrv.Verify(uuid)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/agent/verified", http.StatusSeeOther)
	}
}
