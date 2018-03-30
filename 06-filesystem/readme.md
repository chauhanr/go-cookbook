# Discovering the Filesystem

### Recipe 1 Getting file information
This is a very simple recipe as we have to use the package os.File.Stat() method. This will return FileInfo
which has all the information you need.

### Recipe 2 Creating Temporary files
There are packages in the ioutil that help us create temp file and temp directories

```
    tFile, err := iotuil.TempFile("", "gotcookbook")
    if err != nil{
        panic(err)
    }
    // call is responsible for handling the clean up
    defer os.Remove(t.File.Name())
    fmt.Println("File Name: ",tFile.Name())

    tDir, err := ioutil.TempDir("", "gotcookbookDir")
    if err != nil{
        panic(err)
    }
    defer os.Remove(tDir)
    fmt.Println("Temp Dir: ", tDir.Name())
```
The code above returns a temp file which can be used. but it is developers responsibility to get ride of it.

### Recipe 3 Writing to a file
There are various ways to write content to a file:
* f.WriteString(str) - writes the data to the file.
* io.Copy(f, str)

```
    f, err := os.Create("sample.file")
    ...
    _, err = f.WriteString("Go is awesome!")
    ...
    _,err = io.Copy(f, strings.NewReader("Yeah! Go is awesome!")
    ...
    
```

### Recipe 4 Writing to files from multiple go routines.
The best way to do this is to have a Writer with a mutex attached. Therefore we ned to make a struct that
will have a mutex attached with the writer.

```
type SyncWriter struct{
    m sync.Mutex
    Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (int, error) {
    w.m.Lock()
    defer w.m.Unlock()
    return w.Writer.Write(b)
}
```

The structure above implements the writer interface and overrides the Write method where we get the lock to the writer
and then writes the data and unlocks mutex at the end.

```
    f, err := os.Create("sample.file")
    ...
    wr := &SyncWriter(sync.Mutex{}, f)
    wg := sync.WaitGroup{}
    for _, val := range data {
        wg.Add(1)
        go func(greeting string){
            fmt.Fprintln(wr, greeting)
            wg.Done()
        }(val)
    }
    wg.Wait()
```

the code above makes as many sub routines as the elements of the array and all of them try to write to the file.
Because of the wait group is used to avoid the dead lock problem with go routines.


### Recipe 5 Listing of directory
There a couple of ways we can list the directories:

* **ioutil.ReadDir(path)** - this returns a list of all folders and files in the path.
* **filepath.Walk(path, func(wPath, info os.FileInfo, err))** - this takes a function taht will be called for all directories
that are encountered.

### Recipe 6 Changing file permissions
The os.FileInfo and os.File package can be used to change the stats

```
    f, err := os.Create("/path/to/file")
    if err != nil{
        panic(err)
    }
    defer f.Close()
    fi, err := f.Stat() // get the file info
    ...
    err = f.Chmod(0777)
    ...
    fi, err = f.Stat()

    fmt.Printf(""File permission: %v\n", fi.Mode())
```

### Recipe 7 Creating files and directories
There are various methods and packages that can be used for this purpose:

* os.Create("path/to/file") - creates a file with default permissions.
* os.OpenFile("path/to/file", flag, permission)
    * flag - denoting CREATE, APPEND properties
    * permission- permissions for the file 0777, 0666 etc.
* os.Mkdir("path", permission) - will try to create the folder mentioned in the path if the parent folder
is not present then the method will fail
* os.MkdirAll(path, perm) - this will create all the folders in the path that are not present also

### Recipe 8 Filtering file in a directory
We can filter files and choose them in selective way by using the following api:

* filepath.Glob(pattern)  - matches files with name in the current directory
* filepath.Match()  - also matches the values based on regex patterns.
* filepath.Walk(path, func(){}) the function is the filtering function that decides whether the folder/files
is choosen or not.

