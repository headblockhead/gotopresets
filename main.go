package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		switch strings.ToLower(prompt("Enter command:")) {
		case "store":
			if exit := store(); exit {
				os.Exit(0)
			}
		case "load":
			load()
		}
	}
}

func prompt(msg string) (response string) {
	fmt.Print(msg + " ")
	reader := bufio.NewReader(os.Stdin)
	response, _ = reader.ReadString('\n')
	response = strings.TrimSpace(response)
	return
}

func store() (exit bool) {
	fmt.Println("Store Location")
	name := prompt("Name:")
	loc := prompt("Location:")
	yn := strings.ToLower(prompt("Is this ok? (Y/N):"))

	if yn == "y" {
		fmt.Println("Ok! Storing in presets.txt...")
		f, err := os.OpenFile("./presets.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString(name + "|" + loc); err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func load() {
	req := prompt("Name:")

	file, err := os.Open("./presets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)
		if strings.HasPrefix(line, req) {
			fmt.Println("Match found!")
			total := strings.Split(line, "|")
			fmt.Println("changing", total[1])
			// This won't actually change directory, because `cd` is a shell 'built-in'.
			// I'll put together an example in shell scripting.
			err = os.Chdir(total[1])
			if err != nil {
				panic(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
