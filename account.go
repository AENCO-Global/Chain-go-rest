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