package main

import (
	"os"
	"path/filepath"
	"log"
	"log/syslog"
)

func main(){
	programName := filepath.Base(os.Args[0])
	log.Println("using syslog")
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil{
		log.Fatalf("%s", err)
	}
	sysLog.Crit("Crit: Logging in Go !")
}