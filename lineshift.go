package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
if strings.Contains(string(FileContents),"10 13") {
fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}
}

func main() {
getFileLine("./files/text1.txt")
getFileLine("./files/text2.txt")
}