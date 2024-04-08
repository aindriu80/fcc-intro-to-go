package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		descrpthandle()
	default:
	fmt:
		Println("Run encrypt to encrypt a file, and decrypt to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("file encryption ")
	fmt.Println(" Simple file encryption for your day to day needs.")
	fmt.Println("")
	fmt.Println("Usage:  ")
	fmt.Println(" ")
	fmt.Println("\t go run . encrypt /path/to/your/file ")
	fmt.Println("")
	fmt.Println("Commands:  ")
	fmt.Println("")
	fmt.Println("\t encrypt\t Encrypts a file given a password ")
	fmt.Println("\t decrypt\t Tries to decrypt a file using a password ")
	fmt.Println("\t help\t\t Displays help text ")
	fmt.Println(" ")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		PrintLn("missing the path to the file.  For more information run . help")
		os.Exit(0)

	}
	file := os.Args[2]
	if !validatefile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n File successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		PrintLn("missing the path to the file.  For more information run . help")
		os.Exit(0)

	}
	file := os.Args[2]
	if !validatefile(file) {
		panic("File not found")
	}
	fmt.Print("Enter Password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n File successfully decrypted")
}

func getPassword() []byte {
	fmt.Print("Enter Password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\n Confirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("/n Passwords do not match.  Please try again\n ")
		return getPassword()
	}
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
