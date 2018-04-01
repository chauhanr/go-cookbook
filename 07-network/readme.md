# Connecting to the network

### Recipe 01 - Navigate the internet interfaces
This recipe lists all the interfaces that the local address network has.
We can use the net.Interfaces()  method to get a list of all the interfaces
on your local machine right now.

### Recipe 02 Connect to the remote server
This recipe creates a small web server and then connects to it using the net.Dial()
using the TCP protocol.

### Recipe 03 Resolving domain and ip address
This recipe will try to resolve the domain name to an ip and and ip to domain values.

```
addrs, err := net.LookupAddr("127.0.0.1")
    if err != nil{
        panic(err)
    }
fmt.Println("Resolving the address 127.0.0.1")
    for _,addr := range addrs{
        fmt.Println(addr)
    }
fmt.Println("Resolving the domain localhost")
    ips, err := net.LookupIP("localhost")
    if err != nil{
        panic(err)
    }
    for _, ip := range ips{
        fmt.Println(ip.String())
    }
```

### Recipe 04 Connect to HTTP server using HTTP request
Instead of creating a TCP connection and interacting with the http server this recipe uses the http request

```
    req, err := http.NewRequest(httpMethod, address, params)
    ..
    res, err := http.DefaultClient.Do(req)  // sending request
    data, err := ioutil.ReadAll(res.Body)

    // close the res.Body
    res.Body.Close()
    fmt.Println("Response darta : "+string(data))

```

### Recipe 7 Working with Http Headers
The Http headers are like maps except that the map values can be become array of values and hold more values.

```
header := http.Header{}
header.Set("Auth-X", "abcde2120fg")
header.Add("Auth-X", "defgm34gh45")
fmt.Println(header)

// retrive values
resSlice := header["Auth-X"]
fmt.Println(resSlice)

resFirst := header.Get("Auth-X")
fmt.Printf("The first value for Auth-X is :%s\n", resFirst)

// replace all values of AuthX with new value
header.Set("Auth-X", "newvalue")
fmt.Println(header)

// remove header
header.Del("Auth-X")
fmt.Println(header)
```

### Recipe 8 Redirecting HTTP requests
This recipe gives us the tools needed to handle the redirect requests using golang. The important piece here is the
http.Client property of ** CheckRedirect which takes a func(req *http.Request,via []*http.Reqest) error{} **  where you can
define what needs to be done with the redirect requests and how many times to retry.

### Recipe 9 Sending email using Go

```
func main(){
    var email string
    fmt.Println("Enter your username for smtp (inclusing @gmail.com): ")
    fmt.ScanIn(&email)
    var pass string
    fmt.Println("Enter your account password: ")
    fmt.ScanIn(&pass)

    auth := smtp.PlainAuth("", email, pass, "smtp.gmail.com")

    c, err := smtp.Dial("smtp.gmail.com:587")
    ...
    defer c.Close()
    config := &tls.Config{ServerName: "smtp.gmail.com"}

    if err = c.StartTLS(config); err != nil{
        panic(err)
    }

    if err = c.Auth(auth); err != nil {
        panic(err)
    }
    if err = c.Mail(email); err != nil {
            panic(err)
    }
    if err = c.Rcpt(email); err != nil {
            panic(err)
    }

    w, err := c.Data()
    if err != nil{
        panci(err)
    }
    msg := []byte("Hello this is content")
    if _, err := w.Write(msg); err != nil{
        panic(err)
    }

    err = w.Close()
    if err != nil {
        panic(err)
    }
    err = c.Quit()
    if err != nil {
        panic(err)
    }

}
```
