package main

import (
    "fmt"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/infrastructure"
    "github.com/AENCO-Global/Chain-go-sdk/sdk/utils"
    "net/http"
)

func networkGet(w http.ResponseWriter, r *http.Request) {
    Server := conf.Server  // "api-1.aencoin.io:3000"

    client := infrastructure.NewAPIClient()
    client.ChangeServer(Server)

    a, err := client.NetworkRoutesApi.GetNetworkType()
    if err != nil {
        fmt.Println(utils.Struc2Json(a))
    }
    fmt.Println( utils.Struc2Json(a))
}