# Working with databases

### Recipe 0 Starting a docker container for postgres
postgres can be pulled from docker hub and started using the following commands:

```
$> docker pull postgres:9.6

$> docker run -it --name postgres-dev -p 5432:5432 -e POSTGRES_PASSWORD=admin -d postgres:latest
```

install psql

```
# installing the postgres client
$> sudo apt-get install postgresql-client

# connecting to the database
$> psql -h localhost -p 5432 -U postgres

# ctl+d will disconnect the psql client
```

### Recipe 2 Verify the connection to data base

```
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"context"
)

func main(){
	connStr := "postgres://postgres:postgres@localhost:5432/examples?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Ping O.K")
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)
	err = db.PingContext(ctx)

	if err != nil{
		fmt.Println("Error: "+err.Error())
	}
	conn,err := db.Conn(context.Background())
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	err = conn.PingContext(context.Background())
	if err != nil{
		panic(err)
	}
	fmt.Println("connection Ping O.K")
}
```

* We connect to the data base using the connection string that tries to connect to the examples database.
* db.Ping() - only connects to the database server and pings for its availablility it does not guarantee a connection from the pool.
* db.PingContext() - this method will create a db connection from the pool is used.


### Recipe 3 executing sql statements
This recipe will execute sql statements which will truncate a table and insert statements to the table.

### Recipe 4 prepared statements
Prepared statement is similar to a sql query but in this case the prepared statement is compiled on the database side and then
the client just passes variables that need to be passed and it is executed many times.

### Recipe 5 cancelling query
if you have a situation where you want the query to run in a particular amount of time and if it does not then you need to cancel it
this recipe helps the key here is to use the XXXContext() method in the db package

```
    ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Microseconds)
    rows, err := db.QueryContext(ctx, query)
    // habdle error
    cancelFunc()   // this will block till either the query runs or time out takes place.

    // then handle the row sets for processing.
```

### Recipe 6 Reading result set metadata
This recipe will help us read the meta data about the table that we are querying.

### Recipe 7 Getting data from result sets
A more involved example of getting values from the database. Inspecting the query result set for more data and putting data to structs.


### Recipe 8 Parsing result sets to maps

```
   func parseToMap(rows *sql.Rows, cols []string) map[string]interface{}{
     values := make([]interface{}, len(cols))
     pointers := make([]interface{}, len(cols))

     for i := range values{
        pointers[i] = &values[i]
     }
     if err := rows.Scan(pointers...); err != nil{
        panic(err)
     }

     m := make(map[string]interface{})
     for i := colName := range cols{
        if values[i] == nil{
            m[colName]= nil
        }else{
            m[colName] = values[i]
        }
     }
     return m
   }

```

### Recipe 9 Handling transactions
This shows us how to work with transactions in database.

