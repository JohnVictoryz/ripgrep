package main

import (
   "fmt"
   "os"
   "log"
   "net/http"
   "net/url"
   "os"
   "strings"
   "bufio"
   "io/ioutil"
)

var (
    arg3    string
    arg2    string
)
func download(){
    // Build arg3 from fullPath
    fileURL, err := url.Parse(arg2)
    if err != nil {
        log.Fatal(err)
        fmt.Print("No Url Provided")
    }
    path := fileURL.Path
    segments := strings.Split(path, "/")
    arg3 = segments[len(segments)-1]
 
    // Create blank file
    file, err := os.Create(arg3)
    if err != nil {
        log.Fatal(err)
        fmt.Println("No arg3 Provided")
    }
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }
    // Put content on file
    resp, err := client.Get(arg2)
    if err != nil {
        log.Fatal(err)
        fmt.Print("Error Retriving File")
    }
    defer resp.Body.Close()
 
    size, err := io.Copy(file, resp.Body)
 
    defer file.Close()
 
    fmt.Printf("Downloaded a file %s with size %d", arg3, size)
}
func help(){
     fmt.Print("The GET Command")
     fmt.Println("List Of Commands\n--help for help\n--download Download a file from the World Wide Web(LOL)\n--file do things with files NOT READY YET")
}
func files(){
       if arg2 == '--view'{
           if arg3 == nil {
               fmt.Println("No file name to read")
               log.Fatal("NoFileProv")
           }
           content, err := ioutil.ReadFile(arg3)

           if err != nil {
               log.Fatal(err)
           }

           fmt.Println(string(content))
       } else if arg2 == '--delete' {
           if arg3 == nil {
               fmt.Println("No file name provided to delete")
               log.Fatal("NoFileProv")
           }

           err := os.Remove(arg3)

           if err != nil {
               log.Fatal(err)
           }

           fmt.Println("File", arg3, "Deleted Succesfully")
       } else if arg2 == '--new' {
           if arg3 == nil{
               fmt.Println("No file name provided to create")
               log.Fatal("NoFileProv")
           }
           file, err := os.Create(arg3)

           defer file.Close()

           if err != nil {
               log.Fatal(err)
           }

           fmt.Println("File", arg3, "Created Succesfully")
       } else if arg2 == '--size' {
           if arg3 == nil {
               fmt.Println("No file name provided to check file size")
               log.Fatal("NoFileProv")
           }


           fInfo, err := os.Stat(arg3)

           if err != nil {

               log.Fatal(err)
           }

           fsize := fInfo.Size()

           fmt.Printf("The file size is %d bytes\n", fsize)
       } else if arg2 == '--lastmod'{
           if arg3 == nil{
               fmt.Println("No file name provided to check modified date")
               log.Fatal("NoFileProv")
           }
           fileName := arg3

           fileInfo, err := os.Stat(fileName)

           if err != nil {
               log.Fatal(err)
           }

           mTime := fileInfo.ModTime()

           fmt.Println(mTime)
       }
}

func main() {
    args = os.Args[1]
    arg3 = os.Args[3]
    arg2 = os.Args[2]
      if args == '--download' {
        download()
      } else if args == '--help' {
        help()
      } else if args == '--file'{
        files()
      } else if args == nil{
      fmt.Println("No Argument Providied use --help for help")
      log.Fatal("NoArgumentProvided")
      } else {
      fmt.Println("Invalid Argument use --help for help and list of commands")
      log.Fatal("Invcmd")
      }

}

// Reminder Not Ready still doing bad at go
// Also Contribute if you want
