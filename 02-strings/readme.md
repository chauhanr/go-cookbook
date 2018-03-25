# Strings and Things
These are recipes that will deal with the standard lib in go to handle strings.

### Recipe 1 - Finding substrings in a string
This recipe helps you search if a word or phrase is present in the string.
There are 3 functions that we can use:
* **strings.contains(source, lookFor)** - this will test if the lookFor string is present in the source
* **strings.HasPrefix(source, prefix)** - if the source string starts with prefix value
* **strings.HasSuffix(source, suffix)** - if the source string ends with suffix value

The searches are case sensitive.

### Recipe 2 - Breaking Strings into words [important]
There are again 4 ways we can do this:
1. use the **strings.Field(source)** method that will split the string based on whitespaces.
2. use the **strings.Splut(source, delimiter)** - this will split the string based on the delimiter passed
3. use the **stings.Field(source, func custome{})**  here we have a custome function to split the string.
4. use the **string.regexp.MatchCompile(reg).Split(source, -1)

There are other SplitXXX() methods that help split as well.

### Recipe 3 - Joining Strings
The strings.Join() method has many flavors
1. strings.Join(str []string, value string) - joins all the values in the array by putting value in between them.



