package main

type Config struct {
    Server string
}

func getGlobals () Config {
    var config Config
    /*  Initialise variables and constants here, it may be necessary to functionally
        get values such as from a list of servers.
        1: Ensure the value is in the Struct above
        2: Assign the value below
     */
    config.Server = "api-1.aencoin.io:3000"
    return config
}