package scratch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Printf("%s ", prompt)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}

const PROMPT = "~> "

func RunConsole() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`
	______  ___   _____ _____ _   _ 
	| ___ \/ _ \ /  ___|_   _| \ | |
	| |_/ / /_\ \\ '--.  | | |  \| |
	| ___ \  _  | '--. \ | | | . ' |
	| |_/ / | | |/\__/ /_| |_| |\  |
	\____/\_| |_/\____/ \___/\_| \_/
	`)
	prompt := func() string {
		return GetInput(PROMPT, reader)
	}
	for {
		input := prompt()
		// TODO: How to manually pass the input into Cobra CLI?
		fmt.Println(input)
	}
}
