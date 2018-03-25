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


### Recipe 4 - Concatenating string using buffered writer
Here we use the bytes.Buffer or plain []byte array to concat value of a series of strings.

```
    buffer := bytes.Buffer{}
    // this give a buffer of bytes that can be used to write too
    buffer.WriteString("heelo") // this will allow to write strings to the buffer.


   // the bytes solution is bit more involved and we use the copy function
   buf := make([]byte, 100)
   c := 0
   for _, val := range stringsCol {
      c += copy(buf[c:],[]byte(val))
   }

```

Benchmarking of the 2 scenarios of the following scenarios:
1. plain concatenation using str += str2
2. using bytes.Buffer WriteString()
3. using copy function using []byte

lead to the conclusion that the copy method is the most performant and is marginally better than the WriteString() of bytes.Buffer

However using the + operator to concatenate strings is the worst by a mile.

### Recipe 5 - Align test using tabwriter
The Tab writer is util in the text package of golang. It gives a writer which can be used to print text to the output stream

```
    w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
    fmt.Fprintln(w, "username \t firstname \t lastname")
    fmt.Fprintln(w, "chauhr \t ritesh \t chauhan")
    fmt.Fprintln(w, "asinghal \t ankur \t singhal")
    w.Flush()
```

### Recipe 6 Replacing a sub string within a string
Here we can use the following alternatives:
1. strings.Replace() that will replace one substring at a time.
2. string.NewReplacer(...) can support multiple replacements at a time
3. regexp package that can help replace more complex matching and replacements.


### Recipe 7 Finding substring in text using regular expression package
Regexp package is a very powerful library that helps in getting all or one matching substring using regular expression.
This is particularly useful when you are trying to extract information from a text data.


### Recipe 8 Decoding and encoding string to unicode.
There is an inbuilt packages in golang that help with unicode support - golang.org/x/text/encoding/charmap




