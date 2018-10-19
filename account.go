package main

import (
    "encoding/json"
    "fmt"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/model/account"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/model/blockchain"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/utils"
    "github.com/gorilla/mux"
    "net/http"
)


type Error struct {
    resp int
    status string
    message string
}


func accountGenerate(w http.ResponseWriter, r *http.Request)  {
    type KeyPair struct {
        Priv string `json:"priv"`
        Pub string `json:"pub"`
    }
    params := mux.Vars(r)

    if (1 == len(params) && 64 == len(params["privateKey"]) ) || (0 == len(params)) { //Ensure there is a parameter and its the right length or Zero
        keypair, _ := account.CreateFromPrivateKey(params["privateKey"],blockchain.NetworkType.TEST_NET)
        messageA := KeyPair{Priv: utils.Bt2Hex(keypair.KeyPair.PrivateKey ) ,Pub: utils.Bt2Hex(keypair.KeyPair.PublicKey) }
        json.NewEncoder(w).Encode(messageA)

    } else {
        retVals := Error{ 400, "Bad_Request","Received an invalid Private Key!, correct the error or try no Private Key to generate a new one, which is a good idea even to see what the valid one looks like." }
        fmt.Fprint(w, retVals)
    }

    w.Header().Set("Content-Type","application/json")
}

func accountTransactionsGet(w http.ResponseWriter, r *http.Request) {
    type accountDetail struct {
        Meta string `json:"meta"`
        Account []string `json:"account"`
    }

    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Public Key: %v\n", params["publicKey"])
}

func accountTransactionsIncomingGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Public Key: %v\n", params["publicKey"])
}

func accountTransactionsOutgoingGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Public Key: %v\n", params["publicKey"])
}

func accountTransactionsUnconfirmedGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Public Key: %v\n", params["publicKey"])
}

func accountTransactionsPartialGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Public Key: %v\n", params["publicKey"])
}

func accountMultisigGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Account ID: %v\n", params["accountId"])
}

func accountMultisigGraphGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Account ID: %v\n", params["accountId"])
}