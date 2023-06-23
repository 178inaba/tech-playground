package main

import (
	"github.com/178inaba/tech-playground/go-plugin-example/common"
	"github.com/hashicorp/go-plugin"
)

type Greeter struct{}

func (g Greeter) Greet() (string, error) {
	return "Hello!", nil
}

func main() {
	var greeter Greeter
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"greeter": &common.GreeterPlugin{Impl: greeter},
		},
	})
}
