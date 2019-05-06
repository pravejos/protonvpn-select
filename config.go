package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isInitialized() bool {
	dir := os.Getenv("HOME")
	if len(dir) == 0 {
		debug("$HOME not set.")
		return false
	}
	return exists(dir) && exists(dir+"/.protonvpn-select") && exists(dir+"/.protonvpn-select/.tier")
}

// exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	debug("File/Directory (%v) does not exist.\n", path)
	return false
}

func initialize() bool {
	dir := os.Getenv("HOME")
	if len(dir) == 0 {
		fatal("$HOME not set.")
		return false
	}

	if exists(dir) {
		if !exists(dir + "/.protonvpn-select") {
			err := os.Mkdir(dir+"/.protonvpn-select", 0700)
			if err != nil {
				fatal("Unable to create directory", dir+"/.protonvpn-select")
				return false
			}
		}
	} else {
		fatal("Home directory (%v) does not exist.\n", dir)
		return false
	}

	tier, err := readTier()
	if err != nil {
		fatal("Unable to read tier.")
		return false
	}
	err = writeTierFile(dir+"/.protonvpn-select/.tier", tier)
	if err != nil {
		fatal("Unable to write file.")
		return false
	}
	return true
}

func readTier() (string, error) {
	fmt.Println("ProtonVPN plans:\n1) Free\n2) Basuc\n3) Plus\n4) Visionary")
	fmt.Print("Enter your ProtonVPN plan: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Unable to read input", err)
		return "", err
	}

	switch char {
	case '1':
		char = '0'
		break
	case '2':
		char = '1'
		break
	case '3':
		char = '2'
		break
	case '4':
		char = '3'
		break
	default:
		fmt.Println("Wrong input. Try again.")
		return readTier()
	}

	return string(char), nil
}

func writeTierFile(filePath string, tier string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fatal("Unable to create file", err)
		return err
	}
	l, err := f.WriteString(tier)
	if err != nil {
		fatal("Error while writing file", err)
		f.Close()
		return err
	}
	debug(strconv.Itoa(l) + "bytes written successfully")
	err = f.Close()
	if err != nil {
		fatal("Error while closing file", err)
		return err
	}
	return nil
}
