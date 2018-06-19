package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

type User struct{
	FirstName string
	LastName string
	Age int
	Active bool
}

func main(){
   var buf bytes.Buffer
   enc := gob.NewEncoder(&buf)
   user := User{
   	"Radomir",
   	"Sohlich", 30, true,
   }
   enc.Encode(user)
   fmt.Printf("%X\n", buf.Bytes())

   out := User{}
   dec := gob.NewDecoder(&buf)
   dec.Decode(&out)
   fmt.Println(out.String())

   enc.Encode(user)
   out2 := SimpleUser{}
   dec.Decode(&out2)
   fmt.Println(out2.String())
}

func (u User) String() string{
	return fmt.Sprintf(`{"FirstName": %s, "Last Name":%s, "Age": %d, "Active":%v}`, u.FirstName, u.LastName, u.Age, u.Active)
}

type SimpleUser struct{
	FirstName string
	LastName string
}

func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"FirstName": %s, "Last Name":%s}`, u.FirstName, u.LastName)
}