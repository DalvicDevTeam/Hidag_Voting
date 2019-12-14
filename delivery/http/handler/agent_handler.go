package handler

import (
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent"
)

//AgentHandler handles
type AgentHandler struct {
	tmpl        *template.Template
	categorySrv agent.AgentService
}

// NewAgentHandler creates new agent handler struct
func NewAgentHandler(T *template.Template, catserv agent.AgentService) *AgentHandler {
	return &AgentHandler{tmpl: T, categorySrv: catserv}
}

// AgentIndex handles request on /Agent
func (ah *AgentHandler) AgentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.verify.layout", nil)
	}
}

// AgentLogin handles request on /agent/login
func (ah *AgentHandler) AgentLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.login.layout", nil)
	}
}

// AgentVerified handles request on /agent/verified
func (ah *AgentHandler) AgentVerified(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ah.tmpl.ExecuteTemplate(w, "agent.verified.layout", nil)
	}
}
