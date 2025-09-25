package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("Terjadi error input", err)
			continue
		} else if input == "" || input == "0" {
			fmt.Println()
			fmt.Println("Pilihan tidak valid!")
			fmt.Println()
			continue
		}

		return input

	}
}
