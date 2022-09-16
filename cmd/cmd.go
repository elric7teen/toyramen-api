package cmd

import "toyramen.com/core/cmd/rest"

const (
	local = "localhost:800"
)

func Start() {
	rest.StarHttpService()
}
