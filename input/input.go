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
		} else if input == "" {
			fmt.Println("Input tidak boleh kosong!")
			continue
		}

		return input

	}
}
