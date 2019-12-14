package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/delivery/http/handler"

	_ "github.com/lib/pq"

	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent/repository"
	"github.com/NatnaelBerhanu-1/Hackathon/Hidag_Voting/agent/services"
)

func main() {
	dbconn, err := sql.Open("postgres", "postgres://app_admin:password@localhost/restaurantdb?sslmode=disable")

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

	agentHandler := handler.NewAgentHandler(templ, agentServ)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/agent", agentHandler.AgentIndex)
	http.HandleFunc("/agent/login", agentHandler.AgentLogin)
	http.HandleFunc("/agent/verified", agentHandler.AgentVerified)
	
	http.ListenAndServe(":8080", nil)
}
