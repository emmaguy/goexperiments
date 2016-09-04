package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	gophers := []string{`ʕ◔ϖ◔ʔ`, `\ʕ◔ϖ◔ʔ/`, `ʕﾉ◔ϖ◔ʔノ ︵ ┻━┻`}

	fmt.Println("Please select which cutie you'd like on your clipboard: ")
	for i, gopher := range gophers {
		fmt.Printf("%d. %s\n", (i + 1), gopher)
	}
	fmt.Println("")

	var gopherIndex int
	if _, err := fmt.Scan(&gopherIndex); err != nil {
		fmt.Println("Invalid input:", err)
		os.Exit(1)
	}

	clipboardOutput := gophers[gopherIndex-1]
	clipboard.WriteAll(clipboardOutput)
	fmt.Printf("Wrote \"%s\" to clipboard\n", clipboardOutput)
}
