# Fun with Concurrency

### Recipe 1 Synchronizing access to resources using Mutex
His pattern is very common and for every items that we need to protect with a mutex
we need to make a struct and have mutex as well as data structure also.

so if you have a list that needs to be protected by mutex then the structure will be

```
type SyncList struct{
    m sync.Mutex
    slice []string
}

// every function after this would have a struct like
func (l *SyncList) Append(val interface{}) {
    l.m.Lock()
    defer l.m.UnLock()
    l.slice = append(l.slice, val)
}
```

after this being set up we can use the struct across in go routines without the problem of concurrent access.

### Recipe 2 Creating a map for concurrent access
There are two solutions to this recipe
* Is to use the pattern shown in the recipe 1 were we have a mutex and a map and then wrap all operations of the map
locking and unlocking the mutex before and after an operations respectively.
* Is to use the sync.Map{} package which is inherent routine save.

```
    m := sync.Mutex{}
    wg := &sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i <10; i++{
        go func(idx int){
            m.Store(fmt.Springf("%d", idx), names[idx])
            wg.Done()
        }(i)
    }
    wg.Wait()

    v, ok:= m.Load("1")
    if ok {
        fmt.Printf("For Load key: 1 got %v\n", v)
    }

    v, ok:= m.LoadOrStore("11", "Tim")
    if !ok {
        fmt.Printf("For Load key: 1 got %v\n", v)
    }

    m.Range(func(k, v interface{}) bool{
        key, _ := k.(string)
        t, _ := v.(string)
        fmt.Printf("For index %v got %v\n", key, t)
        return true
    })
```

### Recipe 3 running a code block once

We can use the sync.Once package that has a DO() function that the be used to run the code once.

```
type Source struct{
    m *sync.Mutex
    o *sync.Once
    data []interface{}
}

func (s *Source) Pop() (interface{}, error){
    s.m.Lock()
    defer s.m.Unlock()
    s.o.Do(func(){
        s.data = names
        fmt.Println("Data has been loaded")
    })
    if len(s.data) > 0{
        res := s.data[0]
        s.data = s.data[1:]
        return res, nil
    }
    return nil, fmt.Errorf("No data available")
}
```

### Recipe 4 Pooling resources to be used across go routines.
Pooling can be done using the sync.Pool{} package.

### Recipe 5 choosing the fast response from different resources.
This is the recipe where we have different go routines doing same task like search and we need
to get the task that completes first. We start multiple goroutines and get the first result from
the routine and return that.

### Recipe 6 Propagating the error using the error groups
