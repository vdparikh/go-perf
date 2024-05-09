# Getting Started with PostgreSQL and Go
This doc will walk you through setting up a PostgreSQL database, connecting to it from a Go application, and creating a simple API to write and read data.

The full code can be found in `cmd/db/main.go`

## Prerequisites
Before you begin, ensure you have the following installed:
- Go
- Docker
- VS Code

## Setting Up PostgreSQL
First, you'll need to set up a PostgreSQL database.

1. **Install PostgreSQL**:
   - On macOS: `brew install postgresql`
   - On Windows: Download and install from [PostgreSQL.org](https://www.postgresql.org/download/windows/)
   - You can also run Postgres container if you don't want to install it on your device
     - `docker pull postgres`
     - `docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres`. If you want to use persistant storage than you can add parameter `-v my_dbdata:/var/lib/postgresql/data postgres` to your docker run command
     - If you want to access the database, you can do `docker exec -it some-postgres psql -U postgres`

## Connecting to PostgreSQL from Go
You'll use the `pq` package, which is a pure Go Postgres driver for Go's `database/sql` package.

1. **Install the `pq` Package**:
   ```bash
   go get -u github.com/lib/pq
   ```

2. **Create a Go File**:
   - Create a new file called `main.go`.

3. **Add Code to Connect to PostgreSQL**:
   ```go
   package main

   import (
       "database/sql"
       "fmt"
       "log"

       _ "github.com/lib/pq"
   )

   func main() {
    	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
       db, err := sql.Open("postgres", connStr)
       if err != nil {
           log.Fatal(err)
       }
       defer db.Close()

       err = db.Ping()
       if err != nil {
           log.Fatal(err)
       }

       fmt.Println("Successfully connected!")
   }
   ```

4. **Writing to and Reading from the Database**

    The code uses a user struct as below
    ```go
    type User struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    ```

    1. **Create a Table**:
    - Add the following function to `main.go` to create a table:
        ```go
        func createTable(db *sql.DB) {
            query := `
            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                name TEXT NOT NULL,
                age INT NOT NULL
            );`
            _, err := db.Exec(query)
            if err != nil {
                log.Fatal(err)
            }
        }
        ```

    2. **Insert Data into the Table**:
    - Add a function to insert data:
        ```go
        func insertUser(db *sql.DB, name string, age int) (int, error) {
            sqlStatement := `
            INSERT INTO users (name, age)
            VALUES ($1, $2)
            RETURNING id`
            id := 0
            err := db.QueryRow(sqlStatement, name, age).Scan(&id)
            if err != nil {
                return 0, err
            }
            fmt.Println("New record ID is:", id)
            return id, nil
        }
        ```

    3. **Query Data from the Table**:
    - Add a function to query data:
        ```go
        func getUsers(db *sql.DB) ([]User, error) {
            rows, err := db.Query("SELECT id, name, age FROM users")
            if err != nil {
                return nil, err
            }
            defer rows.Close()

            var users []User
            for rows.Next() {
                var u User
                if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
                    return nil, err
                }
                users = append(users, u)
            }
            return users, nil
        }     
        ```

    4. **Step 4: Exposing an API**
    Use the `net/http` package to create a simple API.
    - Modify your `main` function and add handlers:
        ```go
            func main() {
                db, err := sql.Open("postgres", "user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
                if err != nil {
                    log.Fatal(err)
                }
                defer db.Close()

                createTable(db)

                http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
                    switch r.Method {
                    case "POST":
                        var user User
                        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
                            http.Error(w, err.Error(), http.StatusBadRequest)
                            return
                        }
                        id, err := insertUser(db, user.Name, user.Age)
                        if err != nil {
                            http.Error(w, err.Error(), http.StatusInternalServerError)
                            return
                        }
                        user.ID = id
                        response, err := json.Marshal(user)
                        if err != nil {
                            http.Error(w, err.Error(), http.StatusInternalServerError)
                            return
                        }
                        w.Header().Set("Content-Type", "application/json")
                        w.Write(response)

                    case "GET":
                        users, err := getUsers(db)
                        if err != nil {
                            http.Error(w, err.Error(), http.StatusInternalServerError)
                            return
                        }
                        response, err := json.Marshal(users)
                        if err != nil {
                            http.Error(w, err.Error(), http.StatusInternalServerError)
                            return
                        }
                        w.Header().Set("Content-Type", "application/json")
                        w.Write(response)
                    }
                })

                log.Println("Server starting on :8080...")
                log.Fatal(http.ListenAndServe(":8080", nil))
            }

        ```

## **Run Your Application**:
   - Execute `go run main.go` to start your server.


## **Interact with the APIs**
Now, you can access your API at `localhost:8080/users` using a tool like Postman or cURL to add and retrieve users.

```bash
curl -X POST http://localhost:8080/users \
     -H 'Content-Type: application/json' \
     -d '{"name": "John Doe", "age": 30}'
```