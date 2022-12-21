package main

import (
   "fmt"
   "os"
    "log"
   "net/http"
   "net/url"
   "os"
   "strings"   
)

var (
    fileName    string
    fullURLFile string
)
func download(){
    // Build fileName from fullPath
    fileURL, err := url.Parse(fullURLFile)
    if err != nil {
        log.Fatal(err)
        fmt.Print("No Url Provided")
    }
    path := fileURL.Path
    segments := strings.Split(path, "/")
    fileName = segments[len(segments)-1]
 
    // Create blank file
    file, err := os.Create(fileName)
    if err != nil {
        log.Fatal(err)
        fmt.Println("No Filename Provided")
    }
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }
    // Put content on file
    resp, err := client.Get(fullURLFile)
    if err != nil {
        log.Fatal(err)
        fmt.Print("Error Retriving File")
    }
    defer resp.Body.Close()
 
    size, err := io.Copy(file, resp.Body)
 
    defer file.Close()
 
    fmt.Printf("Downloaded a file %s with size %d", fileName, size)
}
func help(){
     fmt.Print("The GET Command")
     fmt.Println("List Of Commands\n--help for help\n--download Download a file from the World Wide Web(LOL)\n--file do things with files NOT READY YET")
}

func main() {
      args = os.Args[1]
      fileName = os.Args[3]
      fullURLFile = os.Args[2]
      if args == --download {
      download()
      } else if args == --help {
      help()
      } else if args == nil{
      fmt.Println("No Argument Providied use --help for help")
      } else {
      fmt.Println("Invalid Argument use --help for help and list of commands")
      }

}

// Reminder Not Ready still doing bad at go
// Also Contribute if you want
