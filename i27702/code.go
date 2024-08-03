package main

import (
	"strings"

	"golang.org/x/net/html"
)

func main() {
	r := strings.NewReader("<math><template><mo><template>")
	html.Parse(r)
}
