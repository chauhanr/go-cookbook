# Files and Directories

When dealing with files and directories there are 2 packages that are very useful.
* os - this helps perform operations like move, delete, rename the files and dir.
* io - this package allows for reading and writing to the files and directories.
* filepath - this has utility functions like Walk that allow you to traverse the tree.

**flag package** in goland has been created to capture command line arguments as well as creating shell like programs.

### Recipe 01 - Dealing with directories

**Symbolic links** - symbolic links are pointers to files or directories and resolved at the time of access.

Following command are being implemented :

* **pwd** - utility to get the current working directory. Simple program that gets you the current working directory.
* **which** - looks for an installed program in the os. - iterates throught %PATH and then tries to find utility functions.
* **rm** - deleting a file in the file system - os.Remove()
* **mv** - moving or renaming files.
* **traverse** - this program will traverse through the directory specified


## os package

* **Chtimes** - Chtimes(name string, accesstime, modifiedTime time.Time) error  - this will change the access time and modification time for the file.
* **Clearenv** - this clears all the environment variables.
* **Environ** - Environ() []string - all entries are in form "key=value"
* **Executable** - returns the path name for the executable that started the current process.
* **Exit** - Exit(code int)  - causes the current programm to exit using the code given.
* **Expand** - Expand(string, func(string) string) string - this replaces ${var} or $var in the string based on the mapping func.
* **ExpandEnv** - replaces he environment varible values in $var or ${var}
```
    fmt.Println(os.ExpandEnv("$USER lives in ${HOME}"))
```
* **Getenv** - Getenv(value string) string - returns the environment variable value that is set.
* **Geteuid** - returns the numeric user id of the caller
* **Getegid** - returns the numeric group id of the caller.
* **Getgroups** Getgroups() ([]int, error) - array of groups that the caller is a part of.
* **Getpagesize** - returns the underlying system page size.
* **Getpid** - Getpid() int - returns the process id of the caller.
* **Getppid** - gets the parent pid for the caller.
* **Getuid** - return the user id of the caller.
* **Getwd** - Getwd() (dir string, err error) - returns a rooted path name of the current directory
* **Hostname** - Hostname() (string, error) - returns the host name of as returned by the kernel.
* **IsExist** IsExist(error) bool - returns a true or false based on the error which indicates whether the file already exisits.
* **IsPathSeparator** - IsPathSeparator(c uint8) bool
* **IsPermission()** - IsPermission(error) bool - returns a bool based on whether the error indicates permission is denied or not.
* **IsTimeOut(error) bool** - similar ot the isPermission and IsExist.
* **Link** - Link(oldName, newName string) error - creates a newName hardlink to the oldName file.
* **LookupEnv** - LookupEnv(string) (string, bool) - returns the value of the environment variable you are looking for but also returns boolean to
indicate if the variable was found or not.
* **Mkdir & Mkdir** - helps create all directories
* **ReadLink** - ReadLink(string) (string, error) - if the file is a symbolic link then we can get the parth to the real file form the ReadLink()
* **Remove & RemoveAll** - removed a single file or directory, where as remove all removes all files and children entires.
* **Rename** - renames (moves) oldpath to new path. if new path already exists and is not a direcotry, rename replaces it.
* **Symlink** - Symlink(oldName, newName string) error - create a symbolic link of new name as a symbolic link to old name.
* **TempDir** - returns the path to the temp directory to store temp files.

#### os.File
* **Create** - Create(name string) (*File, error) - this creates a file with file permission 0666, truncating the file if it exists.
* **NewFile** - NewFile(fd uintptr, name string) *File - it returns a file with the file descriptor specified.
* **Open & OpenFile** - Open(name string) (*File, error) - opens a said file for reading, the mode to read O_RDONLY
* **(f File) Stat()** - returns a fileInfo structure describing the file.
* **(f File) Sync()** - stores the current content of the file to stable storage.

#### os.FileInfo
```
type FileInfo interface{
    Name() string    // base name of the file
    Size() int64     // length in bytes for regular files;
    Mode() FileMode  // file mode in bits
    ModTime() time.Time // modification time
    IsDir() bool     // true if the file is a directory
    Sys() interface{} // underlying data source (can return nil)
}
```
This is retuned from Lstat() and Stat() functions.

#### os.FileMode
It is of type FileMode uint32  - this represents the file mode and permission bits. You can determine the type of file using model

```
    switch mode := fi.Mode(); {
        case mode.IsRegular():
            fmt.Println("regular file")
        case mode.IsDir():
            fmt.Println("directory")
        case mode&os.ModeSymlink != 0:
            fmt.Println("symbolic link")
        case mode&os.ModeNamePipe != 0:
            fmt.Println("named pipe")
    }
```

#### os.Process

Process struct
```
type Process struct{
    Pid int
}
```

* **FindProcess** - FindProcess(pid int) (*Process, error) - find process given the pid of the process.
* **StartProcess** - StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) start the process using the
attr and arguments to the process. if the calling routine uses the runtime.LOCKOSThread and modifies any inheritable OS Level thread state. then the child process inherits caller thread's state.
* **(p Process) Kill** - kills the process and exits.
* **(p Process) Signal** - this sends a signal code int to the OS.
* **(p Process) Wait**  - returns the ProcessState

#### os.ProcessState
this structure is returned when we call a wait method on the process.

* **(p ProcessState) Exited()** - returns a bool indicating if the process has exited.
* **(p ProcessState) Pid()**
* **(p ProcessState) String()**
* **(p ProcessState) Success()**
* **(p ProcessState) UserTime()** - user CPU time of exited process.


## path/filepath package

* **Abs** - Abs(path string) (string, error) - returns the absolute path of the file system.
* **Base** - Base(string)string - returns the last element of the path. if the path is empty the function returns a "."
* **Glob** - Glob(pattern string) (matches []string, err error) - Return all files matching the pattern passed. the pattern is a regular expression. The error is only sent when the pattern passed is a bad patten.
* **Join** - Join(element ...string) string - combines all the elements given into a path.

```
    // all of the following return the same value a/b/c
    filepath.Join("a", "b", "c")
    filepath.Join("a", "b/c")
    filepath.Join("a/b", "c")
```

* **Split** Split(path string) (dir, file string) - splits the path immediately following the last separator

```
    func main() {
    	paths := []string{
    		"/home/arnie/amelia.jpg",
    		"/mnt/photos/",
    		"rabbit.jpg",
    		"/usr/local//go",
    	}
    	fmt.Println("On Unix:")
    	for _, p := range paths {
    		dir, file := filepath.Split(p)
    		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
    	}
    }
```

```
On Unix:
input: "/home/arnie/amelia.jpg"
        dir: "/home/arnie/"
        file: "amelia.jpg"
input: "/mnt/photos/"
        dir: "/mnt/photos/"
        file: ""
input: "rabbit.jpg"
        dir: ""
        file: "rabiit.jpg"
input: "/usr/local//go"
        dir: "/usr/local//"
        file: "go"
```

* **SplitList** - this splits the paths joined by a OS specific list separator, useful when using $PATH
```
    func main() {
    	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
    }

    returns
    On Unix: ["a/b/c", "/usr/bin"]
```

