package main

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string // script signiture
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
