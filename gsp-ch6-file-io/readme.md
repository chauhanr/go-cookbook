# File Input and Output

#### Recipe 1 - io and bufio

The io and bufio packages have utility functions to read and write to files.

io package code snippet
```
    byteSlice := []byte{"Ritesh Chauhan \n"}
    ioutil.WriteFile(filename, byteslice, 066)

    f, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()
    byteslice2 := make([]byte,100)
    n, err := f.Read(byteslice2)
    fmt.Printf("Read %d bytes: %s", n, byteslice2)

```

bufio - buffered io which is used to read from a file in chunks.

```
    f, err := os.Open(filename)
    if err != nil{
        fmt.Printf("error opening %s: %s", filename, err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan(){
        line := scanner.Text()
        if scanner.Err() != nil{
            fmt.Printf("error reading file %s", err)
            os.Exit(1)
        }
        fmt.Println(line)
    }
```


#### Recipe 2 - copying files in golang

There are many way to copy files in go but we will have a look at to main variations.

1. Reading the file in full and writing to the destination
    * Using the io.Copy(dst, source)
    * Using ioutil package's ReadFile() and WriteFile() function.
2. Reading the file in buffer of fixed size and then writing the file in parts.

code snippets below show the two techniques under pattern 1 where as pattern 2 is implemented.

```
  // io.Copy()

  func Copy(source, dest string) (int64, error) {
     sourceFileStat, err := os.Stat(source)
     if err != nil{
        return 0, err
     }
     if !sourceFileState.Mode().IsRegular() {
        return 0, fmt.Errorf("%s is not a regular file", src)
     }

     src, err := os.Open(source)
     if err != nil{
        return 0, err
     }
     defer src.Close()

     destination, err := os.Create(dest)
     if err != nil{
        return 0, err
     }
    defer destination.Close()
    nBytes, err := io.Copy(destination, src)
    return nBytes, nil
  }

```

```
    // ioutil package

     in, err := ioutil.ReadFile(source)
     if err != nil{
        ..
     }

     err = ioutil.WriteFile(destinationm, input, 0644)
     if err != nil{
        ...
     }
```

