package main

type Block struct {
	Timstamp      int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
