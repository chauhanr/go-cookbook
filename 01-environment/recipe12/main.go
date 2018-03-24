package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Client struct{
	consulIP string
	connString string
}

func (c *Client) String() string{
	return fmt.Sprintf("Consul IP: %s, connection string: %s", c.consulIP, c.connString)
}

var defaultClient = Client{consulIP:"locahost:9900", connString: "postgres://localhost:5432"}

/*declaring a config function taht will be passed around */
type ConfigFunc func(opt *Client)

func FromFile(path string) ConfigFunc{
	return func (c *Client){
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)
		fop := struct{
			ConsulIP string  `json:"consul_ip"`
		}{}

		err = decoder.Decode(&fop)
		if err != nil{
			panic(err)
		}

		c.consulIP = fop.ConsulIP
	}
}

func FromEnv() ConfigFunc{
	return func(c *Client){
		connStr, exists := os.LookupEnv("CONN_DB")
		if exists{
			c.connString = connStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client{
	client := defaultClient
	for _, val := range opts{
		val(&client)
	}
	return &client
}

func main(){
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}

