package main

import "github.com/minskylab/covid19"

func main() {
	api, err := covid19.NewAPI()
	if err != nil {
		panic(err)
	}

	if err := api.Run(8080); err != nil {
		panic(err)
	}
}
