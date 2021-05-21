package main

type TXInput struct {
	Ixid      []byte
	Vout      int
	ScriptSig string // script signiture
}
