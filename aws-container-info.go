package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
    "crypto/md5"
    "encoding/hex"
    "html/template"
    "strings"
)

type ColorCombo struct {
    ContainerID		string
    Hostname		string
    ContainerIDColor	string
    HostnameColor	string    
}

func color_from_string(random_string string) string {
    hasher := md5.New()
    hasher.Write([]byte(random_string))
    colorhash := hex.EncodeToString(hasher.Sum(nil))
    color := colorhash[:6]
    return color
}

func get_containerid() string {
    containerid, err := os.Hostname()
    _ = err
    return containerid
}

var containeridcolor = color_from_string(get_containerid())

func get_hostname() string {
//    resp, err := http.Get("http://169.254.169.254/latest/meta-data/hostname")
    resp, err := http.Get("http://sangilak.com/hostname.html")
    defer resp.Body.Close()
    htmlData, err := ioutil.ReadAll(resp.Body)
    _ = err
    htmlDataString := string(htmlData)
    htmlDataStringClean := strings.TrimSpaces(htmlDataString)
}

var hostnamecolor = color_from_string(get_hostname())

func index_handler(w http.ResponseWriter, r *http.Request) {
    t := template.New("some template")
    t, _ = t.ParseFiles("/index.html")
    p := ColorCombo{ContainerID: get_containerid(), Hostname: get_hostname(), ContainerIDColor: containeridcolor, HostnameColor: hostnamecolor}
    fmt.Println("%v", get_hostname())
    fmt.Println("%v", get_containerid())
    fmt.Println("%v", get_hostname())
    fmt.Println("%v", get_containerid())
    fmt.Println("%v", containeridcolor)
    fmt.Println("%v", hostnamecolor)
    fmt.Println("%v", p)
    t.Execute(w, p)
    fmt.Printf("%v", p)
//    fmt.Fprintf(w, t.Execute(w, p))
}

func main() {
    fmt.Println("Running...")
    http.HandleFunc("/", index_handler)
    http.ListenAndServe(":9000", nil)
}
