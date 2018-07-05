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
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
