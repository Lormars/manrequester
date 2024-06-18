package main

import "github.com/lormars/requester/pkg/runner"

func main() {
	config, _ := runner.NewConfig("http://ape-fe-jlb.integ.amazon.com")
	runner.Run(config)
}
