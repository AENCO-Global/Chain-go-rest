package main
// API Details to Match https://www.notion.so/aencoins/API-Interface-1af6611260b04c899f9912e0b12d9344
import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)


const (
    private = "6C07F78D8C932626F6550FB114C26EFAFE2EC40220E44E1EF0180D9FB89A0AF0"
    public  = "8C2C06CCCDDFBC964345C051B3A94906813DCB198BA8A56378DA6ED1D2E99B58"
)

func main() {
    route := mux.NewRouter()

    route.HandleFunc("/version" , version ).Methods("GET")
    route.HandleFunc("/crypto" , accountGenerate ).Methods("GET")
    route.HandleFunc("/crypto/{privateKey}" , accountGenerate ).Methods("GET")
    route.HandleFunc("/account/{accountId}" , accountGet ).Methods("GET")
    route.HandleFunc("/account/{accountId}" , accountPost ).Methods("POST")
    route.HandleFunc("/account/{publicKey}/transactions" , accountTransactionsGet ).Methods("GET")
    route.HandleFunc("/account/{publicKey}/transactions/incoming" , accountTransactionsIncomingGet ).Methods("GET")
    route.HandleFunc("/account/{publicKey}/transactions/outgoing" , accountTransactionsOutgoingGet ).Methods("GET")
    route.HandleFunc("/account/{publicKey}/transactions/unconfirmed" , accountTransactionsUnconfirmedGet ).Methods("GET")
    route.HandleFunc("/account/{publicKey}/transactions/partial" , accountTransactionsPartialGet ).Methods("GET")
    route.HandleFunc("/account/{accountId}/multisig" , accountMultisigGet ).Methods("GET")
    route.HandleFunc("/account/{accountId}/multisig/graph" , accountMultisigGraphGet ).Methods("GET")
    route.HandleFunc("/blocks/{height}" , blocksGet ).Methods("GET")
    route.HandleFunc("/blocks/{height}/limit/{limit}" , blocksLimitGet ).Methods("GET")
    route.HandleFunc("/blocks/{height}/transactions" , blockTransactionsGet ).Methods("GET")
    route.HandleFunc("/chain/height" , chainHeightGet ).Methods("GET")
    route.HandleFunc("/chain/score" , chainScoreGet ).Methods("GET")
    route.HandleFunc("/diagnostic/storage" , diagnosticStorageGet ).Methods("GET")
    route.HandleFunc("/mosaic/{mosaicId}" , mosaicGet ).Methods("GET")
    route.HandleFunc("/mosaic" , mosaicPost ).Methods("POST")
    route.HandleFunc("/namespace/{namespaceId}/mosaics" , namespaceMosaicGet ).Methods("GET")
    route.HandleFunc("/mosaic/names" , mosaicNamesPost ).Methods("POST")
    route.HandleFunc("/namespace/{namespaceId}" , namespaceGet ).Methods("GET")
    route.HandleFunc("/account/{accountId}" , accountNamespacesGet ).Methods("GET")
    route.HandleFunc("/account/namespaces" , accountNamespacesPost ).Methods("POST")
    route.HandleFunc("/namespace/names" , namespaceNamesPost ).Methods("POST")
    route.HandleFunc("/transaction/{transactionId}" , transactionGet ).Methods("GET")
    route.HandleFunc("/transaction" , transactionPost ).Methods("POST")
    route.HandleFunc("/transaction" , transactionPut ).Methods("PUT")
    route.HandleFunc("/transaction/partial" , transactionPartialPut ).Methods("PUT")
    route.HandleFunc("/transaction/cosignature" , transactionCosignaturePut ).Methods("PUT")
    route.HandleFunc("/transaction/{hash}/status" , transactionStatusGet ).Methods("GET")
    route.HandleFunc("/transaction/statuses" , transactionStatusesPost ).Methods("PUT")
    route.HandleFunc("/network" , networkGet ).Methods("GET")
    route.HandleFunc("/node/info" , nodeInfoGet ).Methods("GET")
    route.HandleFunc("/node/time" , nodeTimeGet ).Methods("GET")

    route.NotFoundHandler = http.HandlerFunc(NotFound)     // If the routing is not found, use this as a default.
    http.ListenAndServe(":3001", route)             // Set up the listerner for the incoming requests
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


func version(w http.ResponseWriter, r *http.Request) {
    var responceText = "Currently we have only 1 version"
    responceText += "Version:1"
    fmt.Fprint(w, responceText)
}

func accountGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)                                   //Get the paramters
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Account ID: %v\n", params["accountId"]) //Return the details
}

func accountPost(w http.ResponseWriter, r *http.Request) {
    // Incoming:
    //    {
    //        "addresses": [
    //        "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY",
    //        "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"
    //        ]
    //    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Account Post")
}

func accountTransactionsGet(w http.ResponseWriter, r *http.Request) {
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

func blocksGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Blocks: [%v]\n", params["height"])
}

func blocksLimitGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Height: [%v] Limit : [%v]\n", params["height"], params["limit"])
}

func blockTransactionsGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Blocks: [%v]\n", params["height"])
}

func chainHeightGet(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{ "height": ["lower", "higher" ] }`)
}

func chainScoreGet(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"scoreHigh": ["lower","higher"],"scoreLow": ["lower","higher"]}`)
}

func diagnosticStorageGet(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"numBlocks": 0,"numTransactions": 0,"numAccounts": 0}`)
}

func mosaicGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "mosaicId: %v\n", params["mosaicId"])
}

func mosaicPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //    {
    //        "mosaicIds": [
    //        "d525ad41d95fcf29"
    //        ]
    //    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "mosaic Post")
}

func namespaceMosaicGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "mosaicId: %v\n", params["namespaceId"])
}

func mosaicNamesPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "mosaicIds": [
    //    "d525ad41d95fcf29"
    //    ]
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "mosaic Names Post")
}

func namespaceGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "namespace Id: %v\n", params["namespaceId"])
}

func accountNamespacesGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "namespace Id: %v\n", params["accountId"])
}

func accountNamespacesPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "addresses": [
    //    "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY",
    //        "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"
    //    ]
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Account names spaces Post")
}

func namespaceNamesPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "namespaceIds": [
    //      "84b3552d375ffa4b"
    //      ]
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Namesspaces Names Post")
}

func transactionGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction ID: %v\n", params["transactionId"])
}

func transactionPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "transactionIds": [
    //    "59B8DA0F0CB2720001103922",
    //        "59B8D8E60CB2720001103904"
    //    ]
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Post")
}

func transactionPut(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "payload": "example"
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Put")
}

func transactionPartialPut(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "payload": "example"
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Partial Put")
}

func transactionCosignaturePut(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "payload": "example"
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Cosignature Put")
}

func transactionStatusGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Hash Status: %v\n", params["hash"])
}

func transactionStatusesPost(w http.ResponseWriter, r *http.Request) {
    // Incoming
    //{
    //    "hashes": [
    //    "5AF72FE39C0515E823903918A0BFE9625DD752C7BD63969C8869F25E9D155DE5"
    //    ]
    //}
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Post")
}

func networkGet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Transaction Hash Status: %v\n", params["hash"])
}

func nodeInfoGet(w http.ResponseWriter, r *http.Request) {
//    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Node Info", `{"publicKey": "EB6839C7E6BD0031FDD5F8CB5137E9BC950D7EE7756C46B767E68F3F58C24390", "port": 7900, "networkIdentifier": 144,"version": 0,"roles": 2,"host": "api-node-0","friendlyName": "api-node-0"}`)
}

func nodeTimeGet(w http.ResponseWriter, r *http.Request) {
//    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Node Info", `{ "communicationTimestamps": { "sendTimestamp": [], "receiveTimestamp": [] }}`)
}


