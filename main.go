package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {

	decode := flag.Bool("d", false, "decode base64 to string")
	encode := flag.Bool("e", false, "encode string to base64")
	flag.Parse()
	if !*decode && !*encode {
		fmt.Println("error: please provide a flag -d (decode) or -e (encode)")
		os.Exit(1)
	}

	if len(os.Args[2:]) != 1 {
		fmt.Println("error: please provide an argument to encode or decode")
		os.Exit(1)
	}

	input := os.Args[2:]
	if *decode {
		decoded, err := base64.StdEncoding.DecodeString(input[0])
		if err != nil {
			fmt.Println("error decoding base64:", err)
			os.Exit(1)
		}
		fmt.Println(string(decoded))
		os.Exit(0)
	}

	if *encode {
		encoded := base64.StdEncoding.EncodeToString([]byte(input[0]))
		fmt.Println(encoded)
		os.Exit(0)
	}
}