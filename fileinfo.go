package main
import (
	"os"
	"fmt"
	"flag"
)

func main() {
FileName := flag.String("f", "", "Filnavnet til file som vi skal hente informasjon fra")
flag.Parse()
fmt.Println("Henter informasjon fra file: " + *FileName)
file, err := os.Open(*FileName) 
if err != nil { // Feil = Panic
         panic("Fatal file error")
     }
defer file.Close()
filestats, err := file.Stat()

if err != nil {
panic("Fatal file error")
}
var bytes int64
bytes = filestats.Size()
var kilobytes int64
kilobytes = (bytes / 1024)

var megabytes int64
megabytes = (kilobytes / 1024)

var gigabytes int64
gigabytes = (megabytes / 1024)

fmt.Println("File infomration for: ", filestats.Name())
fmt.Println("Size(Bytes): ", bytes)
fmt.Println("Size(Kibibytes): ", kilobytes)
fmt.Println("Size(Mibibytes): ", megabytes)
fmt.Println("Size(Gibibytes): ", gigabytes)
fmt.Println("Mode: ", filestats.Mode())
fmt.Println("IsDir: ", filestats.IsDir())

filemodes, err := file.FileMode()

}