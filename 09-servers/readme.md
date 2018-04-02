# Server Side (TCP and UDP)

### Recipe 1 - Creating a TCP server

```
    l, err := net.Listen("tcp", ":8080")
        if err != nil{
            panic(err)
        }
    for {
        fmt.Println("Wating for client...")
        conn, err := l.Accept()
        if err != nil{
            panic(err)
        }
        msg, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil{
            panic(err)
        }
        _, err = io.WriteString(conn, "Received: "+string(msg))
        if err != nil{
            panic(err)
        }
        conn.Close()

    }
```

* The first step is to open a server with TCP protocol is to use net package to listen to the port 8080
* The for loop will all for the server to keep listening for a client to join.
* **l.Accept()** method will block until there is a client that connects to the server.
* Once a client connects to the server the bufio will read the data form the client.
* and then the message that is received is written back to the connection so the client will get back what was written.
* connection is closed and then we go for the loop again.

### Recipe 2 - Create a UDP server
The UDP server is similar to the TCP server

```
pc, err := net.ListenPacket("udp", ":7070")
if err != nil{
    log.Fatal(err)
}
defer pc.Close()

buffer := make([]byte, 2048)
fmt.Println("Waiting for client...")

for {
    _, addr, err := pc.ReadFrom(buffer)
    if err != nil{
        rcvMsg := string(buffer)
        fmt.Println("Received: "+ rcvMsg)
        if _, err := pc.WriteTo([]byte("Received: "+ rcvMsg), addr);
        err != nil {
            fmt.Println("error on write: "+ err.Error())
        }
    }else{
        fmt.Println("error: "+err.Error())
    }
}

```


### Recipe 5 Handling the HTTP request
Here we need to create a http server and then use the mux (multiplexer) that will handle the requests

### Recipe 6 Handling Middlware
check out the code to see a pattern on chaining handler funcs


### Recipe 7 Serving static files
Go makes it really easy to make a file server with the use of standard library.

```
    fileServ := http.FileServer(http.Dir("html"))
    fileServ = http.StripPrefix("/html", fileServ)
    http.Handle("/html/", fileServ)
    http.ListenAndServe(":8080", nil)
```

the package http.FileServer() serves the content of the directory provided to it by exposing it over the web.
This is sort of file server.

### Recipe 8 Handling Cookies
This recipes shows how we can set a cookie and then retrive it.

```
    http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        c := &http.Cookie{
            Name: cookieName,
            Value: "Go is awesome",
            Expires: time.Now().Add(time.Hour),
            Domain: localhost
        }
        http.SetCookie(w,c)
        fmt.Fprintln(w, "Cookie is set")
    })

    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        val, err := r.Cookie(cookieName)
        if err != nil{
            fmt.Fprintln(w, "Cookie err: "+ err.Error())
            return
        }
        fmt.Fprintf(w, "Cookie is: %s\n", val.Value)
        fmt.Fprintf(w, "Other Cookies")
        for _, v := range r.Cookies(){
            fmt.Fprintf(w,"%s => %s \n", v.Name, v.Value)
        }
    })

```

### Recipe 9 Gracefully shutting down the server
This recipe will add the capability to serve existing clients but not take any new requests.


