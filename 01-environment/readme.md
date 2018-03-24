# Interacting with the Enviroment
The recipes here show how we can interact with the operating system and play with arguments and flags.
There are recipes to interact with processes as well.

### Recipe 1 - Retrieve the Golang version.
The version of golang in code can be gotten from the runtime package
```
 log.Printf("version %s",runtime.Version())
```

The Version() method prints a const in the source code of the Golang installation.

### Recipe 2 - passing arguments to the program
Simple recipe that gets and prints all the arguments passed to the executable. The first argument of os.Args is always the executable name followed by other parameters passed.

### Recipe 3 - Making use of flags package
The flag package allows for capturing values from the command line sent to the program therefore we can say -flag value

The flag package is an excellent package if you are developing a command line utility to parse the input that the user is giving.

Important points on the flag package are:
1. The flags like Int, StringVar or Var methods must be defined before the Parse() method is called.
2. The Parse() method parses all parameters in os.Args[1:]

Powerful feature of the flag package is that you can parse even custom flag type that conform to a golang data structure in this case it was a an array of strings.
The flag.Var() take an input at interface{}

```
    type ArrayValue []string
    ...
    var arr ArrayValue
    flag.Var(&arr, "array", "Input array to iterate through")
    // golang can parse the custome types as well.
```
Other interesting feature in the flag package is the FlagSet that allows for flag rules to be clubbed in different set and then you can all rules on a particular set.

### Recipe 4 - Setting default value to env variables.
simple program that can set env variables in the environment. Learned about 3 methods
* os.Getenv() - gets the variable by the name provided
* os.LookupEnv() - return a value and bool indicating of the variable was found.
* os.Setenv() - sets the env variable
* os.Unsetenv() - removes the environment variable.

### Recipe 5 - Getting the current working location
The os packages Executable() method is a good way to get the path to the executable.
* os.Executable() give the path as well as the executable name
* if we pass the executable to the path/filepath package we can get the folder  filepath.Dir(ex)
* The pitfall of earlier call is that the path may be a symlink to overcome this we need to use the **filepath.EvalSymlinks(exPath)**

### Recipe 6 - Getting the current PID
The pid can be easily gotten for the proces using the os.Getpid() method. if we have linux environment on which we run the executable then we can run the ps command to get the current pid details too.

### Recipe 7 - Capturing Operating system signals
This program captures the system signals that the OS sends.

### Recipe 8 - Running external processes
The exec package in golang helps us run external commands easily. There are two ways to do it

method 1
```
   prc :=  exec.Command(...)
  err := prc.Run()
   ...
```

method 2
```
    prc := exec.Command(...)
    err := prc.Start()
      /* handle err and other steps */

    prc.Wait()
    ..

```

In the first method we use the **Run()** method and this will block the main go routine that is calling it. However in method 2
we call the **prc.Start()** that give the control back to the main routine while running the process in its own space. Also the **Wait()**
method is used to wait for the process to end and when it does the Wait() cleans all the resources it holds.


### Recipe 9 - Retrieving child process information
The **ProcessState** of a process exec.Command() return type give all the information of the process. But the process state is
only available once the process has completed it job and not before. However you can get the pid for the process before the process has completed its job.


### Recipe 10 - Reading and writing to child processes
This recipe is all about spawning a new process and capturing the std input and output of the procedure and then sending information to and from the process.
In the example we have two process one that is writing to std in and the other which is reading from it.




