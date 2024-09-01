package database

import (
	"os"
	"path/filepath"
)

type State struct {
    Balances map[Account]uint
    txMempool []Tx

    dbFile *os.File
}

func NewStateFromDisk() (*State,error) {
    cwd,err := os.Getwd()
    if err != nil {
        return &State{},nil
    }

    path := filepath.Join(cwd,"database","genesis.json")

    gen,err := loadGenesis(path)
    if err != nil {
        return &State{},nil
    }

    balances := make(map[Account]uint)
    
    for account,balance := range gen.Balances {
        balances[account] = balance
    }
    

    return &State{},nil
}

