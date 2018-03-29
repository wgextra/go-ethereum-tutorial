package main

import (
    "fmt"
    "context"
    "log"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "io/ioutil"
    "encoding/json"
)

func main() {
    bytes,err := ioutil.ReadFile("config.json")
    if err != nil{
        log.Fatalf("File Error!", err)
    }

    var config Config
    if err := json.Unmarshal(bytes,&config);err != nil{
        log.Fatalf("Json Error!", err)
    }
    conn, err := ethclient.Dial(fmt.Sprintf("https://mainnet.infura.io/%s",config.InfuraId))
    if err != nil {
        log.Fatalf("Dial Error!", err)
    }
    ctx := context.Background()
    tx,pending,err := conn.TransactionByHash(ctx,common.HexToHash(config.ContractHash))
    if(!pending){
        fmt.Println(tx)
    }
}

