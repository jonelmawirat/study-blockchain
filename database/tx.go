package database

type Account string

type Tx struct {
    From        Account     `json:"from"`
    To          Account     `json:"to"`
    Value       uint        `json:"value"`
    Data        string      `json:"data"`
}

func (t *Tx) isReward() bool {
    t.Data = "reward"
    return true
}


