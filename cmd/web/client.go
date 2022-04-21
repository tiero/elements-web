package main

import (
	"encoding/json"
	"log"
)

type EmptyRequest struct{}
type GetBlockchainInfoResponse struct {
	Chain                string  `json:"chain"`
	Blocks               int     `json:"blocks"`
	Headers              int     `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	MedianTime           int     `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress"`
	InitialBlockDownload bool    `json:"initialblockdownload"`
	ChainWork            string  `json:"chainwork"`
	SizeOnDisk           int     `json:"size_on_disk"`
	Pruned               bool    `json:"pruned"`
}

type Elements struct {
	client *RPCClient
}

func (e *Elements) getBlockchainInfo() (*GetBlockchainInfoResponse, error) {
	r, err := e.client.call("getblockchaininfo", nil)
	if err = handleError(err, &r); err != nil {
		log.Fatalf("jsonrpc: %v", err)
	}

	var response GetBlockchainInfoResponse
	if err := json.Unmarshal(r.Result, &response); err != nil {
		log.Fatalf("jsonrpc: %v", err)
	}

	return &response, nil
}
