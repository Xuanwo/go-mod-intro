package main

import (
	v24 "github.com/google/go-github/v24/github"
	v25 "github.com/google/go-github/v25/github"
	"golang.org/x/text/width"
)

var (
	_ = v24.Tag{}
	_ = v25.Tag{}
	// Import golang.org/x/text/width to show minimal version select
	_ = width.EastAsianAmbiguous
	// Import  github.com/yunify/qscamel/utils to show auto import
	//_ = utils.Conn{}
)

func main() {
	return
}
