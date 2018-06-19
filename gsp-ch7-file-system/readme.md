# Working with file systems

log package in go will log the details from your program to sys out (console) were as the log/syslog package will log to the system logging service.

### Putting the data at the end of the file
This can easily be done by opening the file and then appending data to it.

```
     f, err := os.OpenFile(filename, os.O_RWR| os.APPEND| os.O_CREATE, 060)
     if err != nil{
        // handle the errr
      }
      defer f.Close()
      fmt.FPrintf(f, "%s\n", message)
```

### System logging
logging facility is like a category used for logging information. the value of the logging facility can be:
auth, authpriv,cron, local0, local1, local2 local3, local4, local5, local6, local7, UUCP, syslog etc. The system logging
configurations are present at
/etc/rsyslog.conf for a linux machine.

```
    programName := filepath.Base(os.Args[0])
    	log.Println("using syslog")
    	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
    	if err != nil{
    		log.Fatalf("%s", err)
    	}
    	sysLog.Crit("Crit: Logging in Go !")
```

