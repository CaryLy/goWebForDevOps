package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

//重启服务
func reLaunch() {
	cmd := exec.Command("sh", "../deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
func firstPage(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "<h1> Hello,this is my  deploy server</h1>")
	io.WriteString(w, "<h1> Hello,this is my  deploy server3</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
