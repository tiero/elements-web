package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type ConnectionDetails struct {
	RpcUser       string
	RpcPass       string
	RpcHost       string
	RpcPort       string
	P2PPort       string
	RemoteRpcHost string
	RemoteP2PHost string
}

type Page struct {
	Title             string
	ConnectionDetails *ConnectionDetails
	BlockchainInfo    *GetBlockchainInfoResponse
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on http://localshot:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	details := detailsFromEnv()
	blockInfo, err := getBlockchainInfo(details)
	if err != nil {
		log.Fatalf("jsonrpc: %v", err)
	}

	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := Page{
		Title:             "Elements Core",
		ConnectionDetails: details,
		BlockchainInfo:    blockInfo,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalln(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func detailsFromEnv() *ConnectionDetails {
	return &ConnectionDetails{
		RpcUser:       os.Getenv("RPC_USER"),
		RpcPass:       os.Getenv("RPC_PASS"),
		RpcHost:       os.Getenv("RPC_HOST"),
		RpcPort:       os.Getenv("RPC_PORT"),
		P2PPort:       os.Getenv("P2P_PORT"),
		RemoteRpcHost: os.Getenv("REMOTE_RPC_HOST"),
		RemoteP2PHost: os.Getenv("REMOTE_P2P_HOST"),
	}
}

func getBlockchainInfo(details *ConnectionDetails) (*GetBlockchainInfoResponse, error) {
	client, err := NewClient(details.RpcHost, details.RpcPort, details.RpcUser, details.RpcPass, false, 30)
	if err != nil {
		return nil, err
	}
	service := &Elements{client}
	response, err := service.getBlockchainInfo()
	if err != nil {
		return nil, err
	}

	return response, nil
}
