package main

import (
	"fmt"

	"github.com/aymerick/raymond"
)

func main() {
	tpl := `<div class="entry">
  <h1>{{title}}</h1>
  <div class="body">
    {{body}}
  </div>
</div>
`

	ctx := map[string]string{
		"title": "My New Post",
		"body":  "This is my first post!",
	}

	result, err := raymond.Render(tpl, ctx)
	if err != nil {
		panic("Please report a bug :)")
	}

	fmt.Print(result)
}
