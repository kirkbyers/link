package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kirkbyers/link"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%+v\n", links)
}
