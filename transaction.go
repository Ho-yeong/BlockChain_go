package main

import "fmt"

const subsidy = 21000 // 보상의 양

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to `%s`", to)
	}

	txin := TXInput{
		Ixid:      []byte{},
		Vout:      -1,
		ScriptSig: data,
	}
	txout := TXOutput{
		Value:        subsidy,
		ScriptPubKey: to,
	}
	tx := Transaction{
		ID:   nil,
		Vin:  []TXInput{txin},
		Vout: []TXOutput{txout},
	}
	tx.SetID()
	return &tx
}

func (tx *Transaction) SetID() {}
