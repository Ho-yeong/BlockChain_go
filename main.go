package main

func main() {
	bc := CreateBlockchain("134")

	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
