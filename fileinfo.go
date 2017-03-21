package main
import (
	"os"
	"fmt"
    "path/filepath"
	"log"
	"flag"
)

func main() {
   dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("Running from: "  + dir)
	
FileName := flag.String("f", "", "Filnavnet til file som vi skal hente informasjon fra")
flag.Parse()
fmt.Println("Henter informasjon fra file: " + *FileName)
file, err := os.Open(*FileName) 
if err != nil { // Feil = Panic
		fmt.Println(err)
         panic("Fatal file error ")
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
fmt.Println("Mode: ", filestats.Mode().String())
fmt.Println("Permission bits: ", filestats.Mode().Perm())
fmt.Println("IsDir: ", filestats.IsDir())
//check hvis Symbolic link
if filestats.Mode()&os.ModeSymlink == os.ModeSymlink {
                 fmt.Println("File is a symbolic link")
}
if filestats.Mode()&os.ModeDevice != 0 {
fmt.Println("Device file.")
}
if filestats.Mode()&os.ModeCharDevice != 0 {
fmt.Println("Unix character device")
}
if filestats.Mode()&os.ModeAppend != 0 {
fmt.Println("Append only device.")
}
}