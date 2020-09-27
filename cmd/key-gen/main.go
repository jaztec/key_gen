package main

import "github.com/jaztec/key_gen"

func main() {
	cfg, _ := key_gen.NewConfig()
	if err := key_gen.PrintKeys(cfg, nil, nil); err != nil {
		panic(err)
	}
}

