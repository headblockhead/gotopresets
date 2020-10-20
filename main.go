package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	loop := true
	for loop == true {
		loop = false
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		if strings.EqualFold(strings.TrimSpace(text), "store") {
			fmt.Println("Store Location")

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Name: ")
			text, _ := reader.ReadString('\n')

			name := text

			reader = bufio.NewReader(os.Stdin)
			fmt.Print("Location: ")
			text, _ = reader.ReadString('\n')

			loc := text

			reader = bufio.NewReader(os.Stdin)
			fmt.Print("Is this ok? (Y/N): ")
			text, _ = reader.ReadString('\n')

			if strings.EqualFold(strings.TrimSpace(text), "Y") {
				fmt.Println("Ok! Storing in presets.txt...")
				if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
					_, err := os.Create("presets.txt")
					if err != nil {
						fmt.Println(err)
						return
					}
				}
				f, err := os.OpenFile("./presets.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					panic(err)
				}

				defer f.Close()

				if _, err = f.WriteString(strings.TrimSpace(name) + "|" + strings.TrimSpace(loc)); err != nil {
					panic(err)
				}
			}
			if strings.EqualFold(strings.TrimSpace(text), "N") {
				loop = true
			}
		} else if strings.EqualFold(strings.TrimSpace(text), "load") {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Name: ")
			req, _ := reader.ReadString('\n')

			file, err := os.Open("presets.txt")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, req) {
					fmt.Println("Match found!")
				}
				total := strings.Split(line, "|")
				_, err = exec.Command("C:/Windows/System32/cmd.exe", "/c" + " cd " + total[1]).Output()
				if err != nil {
					panic(err)
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}

		}
	}
}
