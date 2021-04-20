package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	cmd := os.Args[1:]

	http.HandleFunc("/", createHandler(cmd))

	addr := ":5050"
	fmt.Printf("Starting the server on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func createHandler(cmd []string) http.HandlerFunc {
	if len(cmd) == 0 {
		panic("Command missing.")
	}
	
	name := cmd[0]
	args := cmd[1:]

	return func(res http.ResponseWriter, req *http.Request) {
		c := exec.Command(name, args...)
		output, err := c.Output()

		if err != nil {
			log.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.Header().Add("Content-Type", "text/plain")
		res.Write(output)
	}
}
