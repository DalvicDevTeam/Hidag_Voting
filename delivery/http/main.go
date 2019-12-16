package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/delivery/http/handler"

	_ "github.com/lib/pq"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent/repository"
	voteRepository "github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/votingmachine/repository"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent/services"
	voteService "github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/votingmachine/services"
)

func main() {
	dbconn, err := sql.Open("postgres", "postgres://postgres:password@localhost/hidag?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	templ := template.Must(template.ParseGlob("../../ui/templates/*.html"))

	agentRepo := repository.NewAgentRepositoryImpl(dbconn)
	agentServ := services.NewAgentServiceImpl(agentRepo)

	voterRepo := voteRepository.NewVMRepositoryImpl(dbconn)
	voteServ := voteService.NewVMServiceImpl(voterRepo)

	agentHandler := handler.NewAgentHandler(templ, agentServ)
	voterHandler := handler.NewVMHandler(templ, voteServ)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/voter", voterHandler.VoterIndex)
	http.HandleFunc("/voter/login", voterHandler.VoterLogin)
	http.HandleFunc("/voter/vote", voterHandler.VoterVote)

	http.HandleFunc("/agent", agentHandler.AgentIndex)
	http.HandleFunc("/agent/login", agentHandler.AgentLogin)
	http.HandleFunc("/agent/verifyvoter", agentHandler.AgentVerifyVoter)
	http.HandleFunc("/agent/openpc", agentHandler.Openpc)

	http.ListenAndServe(":8080", nil)
}
