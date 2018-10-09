package main

import (
    "fmt"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/model/account"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/model/blockchain"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/utils"
    "github.com/gorilla/mux"
    "net/http"
)

func accountGenerate(w http.ResponseWriter, r *http.Request)  {
    params := mux.Vars(r)                                   //Get the paramters

    if (1 == len(params) && 64 == len(params["privateKey"]) ) || (0 == len(params)) { //Ensure there is a parameter and its the right length or Zero
        keypair, _ := account.CreateFromPrivateKey(params["privateKey"],blockchain.NetworkType.TEST_NET)
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Private Key: %v\n", utils.Bt2Hex(keypair.KeyPair.PrivateKey )) //Return the details
        fmt.Fprintf(w, "Public Key: %v\n", utils.Bt2Hex(keypair.KeyPair.PublicKey )) //Return the details
        fmt.Fprintf(w, "Incoming  Key: %v\n", params["privateKey"]) //Return the details
    } else {
        fmt.Fprint(w, JsonResponse{"resp": 400, "status": "Bad_Request","message":"Received an invalid Private Key!, correct the error or try no Private Key to generate a new one, which is a good idea even to see what the valid one looks like." } )
    }
}