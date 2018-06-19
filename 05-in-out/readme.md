# Input and Output

### Recipe 1 - Reading from standard input
There are various packages that can help you read form the stdin.

* fmt package Scan()
* bufio.NewScanner()
* os.Stdin.Read()

All are similar is functionality
```
      // using the fmt package
      var name string
      fmt.Println("What is your name? ")
      fmt.Scanf("%s\n", &name)

      var age int
      fmt.Println("What is your age? ")
      fmt.Scanf("%d\n", &age)
      fmt.Printf("Hello %s you age is %d\n", name, age)
```

bufio package
```
    // scanner
    sc := bufio.NewScanner(os.Stdin)

    for sc.Scan(){
        txt := sc.Text()
        fmt.Printf("Echo: %s\n", txt)
    }
```

os.Stdin.Read

```
    for {
        data := make([]byte,8)
        n, err := os.Stdin.Read(data)

        if err == nil && n>0{
            process(data)
        }else{
            break
        }
    }

    ....
    process(data []byte) {}  // could either print or interpret the values.
```


### Recipe 2 - writing to standard output and error

Each process has a stdin, stdout and stderr file descriptors. as these are file discriptors the destination
were the data is written can be anything from console to a socket. There a couple of ways to write to these:

* io.WriteXXX()
* os.Stdout.Write()

```

io.WriteString(os.Stdout, "This is string to standard output.\n")
	io.WriteString(os.Stdout, "This is string to standard error.\n")

	buf := []byte{0xAF, 0xFF, 0xFE}

	for i:=0; i<20; i++{
		if _, e:= os.Stdout.Write(buf); e != nil{
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout, "\n")

```

### Recipe 4 Reading file into a string
First task is to read the file content and then write to a string reader
The follow code snippet shows how we can read from a file and write bufio.NewScanner() to a writer or
we can get content of the file directly using the ioutil package

```
    fmt.Println("### Read as a reader###")
    f, err := os.Open("temp/file.txt")
    if err != nil{
        panic(err)
    }
    defer f.Close()
    // read from the file reader
    wr := bytes.Buffer{}
    sc := bufio.NewScanner(f)
    for sc.Scan(){
        wr.WriteString(sc.Text())
    }
    fmt.Println(wr.String())

    fmt.Println("### Read File ###")
    fContent, err := ioutil.ReadFile("temp/file.txt")
    if err != nil{
        panic(err)
    }
    fmt.Println(string(fContent))

```

### Recipe8 Writing to multiple writer at once
We need to use the io.Multiwriter() for this purpose. The code snippet shows how we can achieve this
```
    buf := bytes.NewBuffer([]byte{})
    f, err := os.OpenFile("sample.txt", os.O_CREATE | os.O_RDWR, os.ModePerm)
    if err != nil{
        panic(err)
    }
    wr := io.MultiWriter(buf, f)  // we have taken two writers and send them to the varidic function.
    _, err = io.WriteString(wr, "Hello, Go is awesome")
    if err != nil{
        panic(err)
    }
    fmt.Println("Content of biffer: "+buf.String())
```

### Recipe 9 Piping between readers and writer.
In Linux operating system the concept of pipes are used extensively. In go we have do the same like we can
pipe data from one socket to another socket. We can do a lot of other interesting stuff.

We can use the io.Pipe to achieve this - io.Pipe is an in-memory pipe and returns both ends of the pipe,
The PipeReader on one end and PipeWriter at the other. Each Write() call on the pipe writer is blocked util
there is a corresponding Read() call on the other.

```
    pReader, pWriter := io.Pipe()

    cmd := exec.Command("echo", "Hello Go!\nThis is an example")
    cmd.Stdout = pWriter

    go func(){
        defer pReader.Close()
        if _, err := io.Copy(os.Stdout, pReader); err != nil{
            log.Fatal(err)
        }
    }()

    if err := cmd.Run(); err != nil{
        log.Fatal(err)
    }
```

here we assign a pWriter to the commands Stdout. When the cmd.Run() is called the commands output is written
to the pWriter pipe. this write is then read by the pReader and subseqently written to the os.StdOut using
the io.Copy(destination pipe, source pipe) functions. Hence the final out put is written to the os.Stdout.

### Recipe 10 Serializing objects to binary format
This is very similar to encoding and decoding json or xml or yaml files. we can also use a binary format that
golang supports or has called gob and the encode and decode funcs are present in the gob package

```
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    // create a user struct with name, age and active attributes.
    enc.Encode(user)  // this step will take the user and move it as bytes.

    dec := gob.NewDecoder(&buf)
    dec.Decode(&out)
    fmt.Println(out.String())
```

### Recipe 11 Handling XML file effectively

```
    type Book struct{
        Title string `xml:"title"`
        Author string  `xml:"author"`
    }

    func main(){
        f, err := os.Open("data.xml")
        if err != nil{
            panic(err)
        }
        defer f.Close()
        decoder := xml.NewDecoder(f)

        books := make([]Book, 0)
        for{
            tok, _:= decoder.Token()
            if tok == nil{
                break
            }
            switch tp := tok.(type) {
                case xml.StartElement:
                    if tp.Name.Local == "book"{
                        var b Book
                        decoder.DecodeElement(&b , &tp)
                        books = append(books,b)
                    }
            }
        }
        fmt.Println(books)
    }
```

