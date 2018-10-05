package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type JsonResponse map[string]interface{}

func NotFound(w http.ResponseWriter, r *http.Request)  {
    var responceText = "Not Found - Invalid Parameters"
    responceText += "The format is dns/{function}/{p1}/{p2}/{etc}</b>"
    responceText += "example:</b> api-1.aencoin..io:3001/version"
    responceText += "will return:</b> <br /> ver:1.(002) "

    fmt.Fprint(w, JsonResponse{"resp": 400, "status": "unknow_parameters","message":responceText } )
}

func (jr JsonResponse) String() (str string) { // method to print out String
    byte, err := json.Marshal(jr)
    if err != nil {
        str = ""                              // return empty
        return
    }
    str = string(byte)                        //ok,return cast byte to string
    return
}


