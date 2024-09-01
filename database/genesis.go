package database

import(
    "os"
    "encoding/json"
)

type genesis struct {
    Balances map[Account]uint `json:"balances"`
} 

func loadGenesis(path string) (genesis,error) {
    contents_in_bytes,err := os.ReadFile(path)
    if err != nil {
        return genesis{},err
    }

    var genesis_data genesis
    err = json.Unmarshal(contents_in_bytes,&genesis_data)

    if err != nil {
        return genesis{},err
    }

    return genesis_data,nil
}

