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

func validateJsonBody(body []byte) string {   // function for validating the Json body
    // validate the key-value of post body json format
    c := make(map[string]interface{})
    json.Unmarshal(body, &c)
    k := make([]string, len(c))

    // copy c's keys into k
    j := 0
    for s, _ := range c {
        k[j] = s
        j++
    }
    var valid_flag int
    var reqst = []string{"version"}

    for _, reqst_val := range reqst {
        for _, k_val := range k {
            if reqst_val == k_val {
                valid_flag++
                break
            }
        }
    }
    if valid_flag == 1 {
        return "ok"
    } else {
        return "not_ok"
    }
}
