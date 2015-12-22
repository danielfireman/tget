package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const sshCmd = "ssh -t -o StrictHostKeyChecking=no"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: tget [remote file to download] [tunnel host 1] ... [tunnel host 2]")
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid remote file to download %s", os.Args[1])
	}
	tunnel := os.Args[2:]
	fURI := os.Args[1]
	fName := path.Join("/tmp", filepath.Base(u.Path))
	var cmd string
	if len(tunnel) > 0 {
		cmd += fmt.Sprintf("%s %s wget %s -O %s", sshCmd, tunnel[len(tunnel)-1], fURI, fName)
	}
	for i := len(tunnel) - 2; i >= 0; i-- {
		cmd = fmt.Sprintf("%s %s '%s; scp %s:%s %s'", sshCmd, tunnel[i], cmd, tunnel[i+1], fName, fName)
	}
	cmd += fmt.Sprintf("; scp %s:%s %s", tunnel[0], fName, fName)
	fmt.Println("Command to execute:", cmd)
}
