package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type ColorCombo struct {
	ContainerID      string
	Hostname         string
	ContainerIDColor string
	HostnameColor    string
}

func perror(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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
	perror(err)
	return containerid
}

var containeridcolor = color_from_string(get_containerid())

func get_hostname() string {
	//    resp, err := http.Get("http://169.254.169.254/latest/meta-data/hostname")
	resp, err := http.Get("http://sangilak.com/hostname.html")
	perror(err)
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	perror(err)
	htmlDataString := string(htmlData)
	return htmlDataString
}

var hostnamecolor = color_from_string(get_hostname())

func index_handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	perror(err)
	p := ColorCombo{ContainerID: get_containerid(), Hostname: get_hostname(), ContainerIDColor: containeridcolor, HostnameColor: hostnamecolor}
	t.Execute(w, p)
}

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":9000", nil)
}
