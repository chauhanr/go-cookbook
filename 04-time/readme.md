# Time and Date

### Recipie 2 Formatting date
time.Format function gives a lot of functionality that can be used to format the date is various ways

```
    tTime := time.Date(2017,time.March,5,8,5,2,0,time.Local)
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	// time for hours and mins
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))
	// predefined date format can be used as well.
	fmt.Printf("The time is: %s\n",tTime.Format(time.RFC1123))
	// zero padding
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))
```

### Recipe 3 Parsing String to date
time.Parse() function is used to parse date to convert the string to time - the parse method give the time is UTC zone if we do not provide the time zone.


```
    t, err := time.Parse("2/1/2006", "31/7/2015")  // Parse(format, dateStr string)
    ...
    fmt.Println(t)
    t, err = time.Parse("2/1/2006 3:23 MST", "31/7/2015 1:25 DST")
    ...
    fmt.Println(t)

```
output will be
* 2015-07-31 00:00:00 +0000 UTC
* 2015-07-31 01:25:00 +0000 DST

### Recipe 4 Converting time to epoch and vice versa
epoch is a universal system to describe the point in time. The begining of epoch time is defines as 00:00:00 1 Jan 1970 UTC.
So getting time since epoch means trying to find the number of seconds or milli seconds that have passed from epoch to the current time.

```
    t := time.Unix(0,0)
    fmt.Println(t)
    // getting the epoch
    epoch := t.Unix()
    fmt.Println(epoch)

    apocNow := time.Now().Unix()
    fmt.Printf("Epoch time in seconds: %d \n", apocNow)

    apocNano := time.Now().UnixNano()
    fmt.Printf("Epoch time in nano seconds: %d\n", apocNano)

```

### Recipe 09 Making piece of code run periodically
This is useful in the case you want a ticker indicator.

### Recipe 11 Timeout long running process
time.After(time) can be used for a time out