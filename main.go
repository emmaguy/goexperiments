package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	var gophers = []string{`ʕ◔ϖ◔ʔ`, `\ʕ◔ϖ◔ʔ/`, `ʕﾉ◔ϖ◔ʔノ ︵ ┻━┻`}

	fmt.Println("Please select which cutie you'd like on your clipboard: ")
	for i, gopher := range gophers {
		fmt.Printf("%d. %s\n", (i + 1), gopher)
	}
	fmt.Println("")

	var gopherIndex int
	fmt.Scan(&gopherIndex)

	clipboardOutput := gophers[gopherIndex-1]
	clipboard.WriteAll(clipboardOutput)
	fmt.Printf("Wrote \"%s\" to clipboard\n", clipboardOutput)
}
