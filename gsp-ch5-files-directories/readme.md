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




