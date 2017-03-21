package main

import (
	"fmt"
	"io/ioutil"
)
var FileContents string
var HexConverted string

func getFileLine(fileName string) {
fmt.Println("Henter Line break format for: " + fileName)
//Lese filen
FileContents, err := ioutil.ReadFile(fileName)
//Hvis error ikke er "nil", panic. 
if err != nil {
panic(err)
}
//Printe ut file-contents
fmt.Println("----------FILE CONTENTS----------")
fmt.Println(FileContents)
fmt.Println("----------FILE CONTENTS----------")

}

func main() {
getFileLine("./files/text2.txt")
getFileLine("./files/text2.txt")
fmt.Println("Main func")
}