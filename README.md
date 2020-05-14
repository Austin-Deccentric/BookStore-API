This is a simple HTTP REST APi that provides a book service.

#### Features:
- Routes POST and GET requests on `api/v1/books`
- Routes PUT and DELETE requests on `api/v1/books/`
- Implemented using MVC architecture.

#### Sample requests:
1. `curl http://localhost:3000/api/v1/books` For a GET request

2. `curl http://localhost:3000/api/v1/books/id` For a DELETE request

3. `curl -H "Content-Type: application/json" -X POST -d {\"id\":1,\"name\":\"Gifted hands\",\"author\":\"Ben Carson\"} http://localhost:3000/api/v1/books` For a POST request(ON windows)

4. `curl -H "Content-Type: application/json" -X POST -d {"id":1,"name":"Gifted hands","author":"Ben Carson"} http://localhost:3000/api/v1/books` For a POST request(For Unix or Mac OSX)

5. `curl -H "Content-Type: application/json" -X PUT -d '{\"name\":\"Gifted hands\",\"author\":\"Ben Carson\"}' http://localhost:3000/api/v1/books/id` For a PUT request(For Windows OS)

6.  `curl -H "Content-Type: application/json" -X POST -d '{"name":"Gifted hands","author":"Ben Carson"}' http://localhost:3000/api/v1/books` For a PUT request(For Unix or Mac OSX)

### Guildline:
* To reproduce you must have MySql server running on your machine.
* Setup a database schema.
* Rewrite line 14 in model/connect.go to:
    * `db,err :=sql.Open("mysql", "\<username\>:\<password\>@tcp(localhost:\<port\>)/<database_name>")`
The variable \<port\> is the port the SQL Server is running on.
