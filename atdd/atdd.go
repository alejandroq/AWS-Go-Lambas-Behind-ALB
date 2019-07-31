package main

import "fmt"

func main() {
	fmt.Println(`
		# To run ATDD Integration Tests
		
		In *Terminal A*, run "make start". This will kickstart a local AWS Lambda RPC on "127.0.0.1:8001".

		In *Terminal B*, run "make integration". This will run a slew of ATDD tests against *Terminal A's* Lambda RPC process.
	`)
}
