package main

import (
	"flag"
	"fmt"
	"os"
	rot "rot13/pkg"
)

func main() {
	mode := flag.String("mode", "", "Describes the mode in which the program executes. Can only be either 'decode' or 'encode'")
	payload := flag.String("payload", "", "the data you want to encode or decode")
	shift := flag.Int("shift", 13, "How many characters should the algorithm skip?")

	flag.Parse()


	if *mode == "" || *payload == "" || *shift > 25 || *shift < 1{
		fmt.Println("rot: missing or invalid operands\nTry 'rot -help' for more information")
		os.Exit(1)
	}

	if *mode == "encode" {
		fmt.Fprintln(os.Stdout, rot.Encode(*payload, *shift))
		return
	}

	fmt.Fprintln(os.Stdout, rot.Decode(*payload, *shift))
}



