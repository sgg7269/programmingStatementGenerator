package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/tjarratt/babble"
)

const (
	appendToStatements = false
	numberOfStatements = 1000000
	cFileName          = "c.statements"
	jsFileName         = "js.statements"
)

func makeCStatement(name, value string) string {
	return fmt.Sprintf("int %s = %s;", name, value)
}

func makeJsStatement(name, value string) string {
	return fmt.Sprintf("var %s = %s;", name, value)
}

func main() {
	var cFile, jsFile *os.File
	var err error

	if appendToStatements == true {
		cFile, err = os.OpenFile(cFileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Couldn't open file: %s\nCreating one instead...\n", cFileName)
			cFile, err = os.Create(cFileName)
			if err != nil {
				fmt.Printf("Couldn't create file: %s\n", cFileName)
				return
			}
		}

		jsFile, err = os.OpenFile(jsFileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Couldn't open file: %s\nCreating one instead...\n", jsFileName)

			jsFile, err = os.Create(jsFileName)
			if err != nil {
				fmt.Printf("Couldn't create file: %s\n", jsFileName)
				return
			}
		}
	} else {
		cFile, err = os.Create(cFileName)
		if err != nil {
			fmt.Printf("Couldn't create file: %s\n", cFileName)
			return
		}

		jsFile, err = os.Create(jsFileName)
		if err != nil {
			fmt.Printf("Couldn't create file: %s\n", jsFileName)
			return
		}
	}

	babbler := babble.NewBabbler()
	babbler.Count = 1

	cWriter := bufio.NewWriter(cFile)
	jsWriter := bufio.NewWriter(jsFile)

	var name, value string
	for i := 1; i < numberOfStatements; i++ {
		rand.Seed(time.Now().UnixNano())
		name = babbler.Babble()
		value = fmt.Sprintf("%d", (rand.Intn(100)))

		_, err = cWriter.WriteString(makeCStatement(name, value) + "\n")
		if err != nil {
			log.Fatal(err)
		}

		_, err = jsWriter.WriteString(makeJsStatement(name, value) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	rand.Seed(time.Now().UnixNano())
	name = babbler.Babble()
	value = fmt.Sprintf("%d", (rand.Intn(100)))

	_, err = cWriter.WriteString(makeCStatement(name, value))
	if err != nil {
		log.Fatal(err)
	}

	_, err = jsWriter.WriteString(makeJsStatement(name, value))
	if err != nil {
		log.Fatal(err)
	}

	err = cWriter.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal(err)
	}

	err = jsWriter.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal(err)
	}

	err = cFile.Close()
	if err != nil {
		fmt.Printf("Couldn't write to file: %s\n", cFileName)
	}

	err = jsFile.Close()
	if err != nil {
		fmt.Printf("Couldn't write to file: %s\n", jsFileName)
	}
}
