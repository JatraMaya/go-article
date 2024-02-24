# README

## How To Run the project
* Extract the file to desired directory
* Run the command `$ go mod download` to download all the necessary dependancies
* Filled up the config.yaml file located at ./config/config.yaml with your desired value e.g. PORT location
* Change to the project root directory and run command `$ go run main.go`
* Use tools like postman to test all the endpoints

## Project specification
- The Project utilize gin as framework for the RestAPI application
- The Project use sqlite3 as database choices due to it's simple and easy to use approach
- Gorm is used as ORM to connect the database with the rest of the project and handle database migration
- The Project utilize jwt token that stored on cookies on successfull login process to protect various endpoints
- The project utilize bcrypt to hash user password to the database so no plain password is stored.

## List of endpoints available
### User related endpoints
- /signup with HTTP method POST require 3 parameters Name, Email, and Password, the Email needs to be unique or it'll received an error
- /login with HTTP method POST authenticate user with 2 parameters Email, and Password, on successfull login process a jwt token will be stored in the cookie for authentication process
- /login with HTTP method GET doesn't require any request body, act as a logout mechanism where it simply delete the existing cookie
### Article related endpoints
- /article with HTTP method GET doesn't require any request body parameters, require valid jwt token presence in cookie to be access
- /article/:id with HTTP method GET, doesn't require any request body parameters, require valid jwt token presence in cookie to be access, utilize id value from url parameters to get specific article with corresponding id
- /article with HTTP method POST, require 3 parameters Title, as in Title of the article, Author, as in Author's name, and Content for the whole content of the article, require valid jwt token to be accessed
- /article/:id with HTTP method PUT, require atleast 1 request body parameter and id value from url parameter to make change to the corressponding article based on key existing in the request body, require jwt token to be access.
- /article/:id with HTTP method DELETE, doesn't require any requst body parameter, require id value from url parameter, will delete article with corresponding id on success, require valid jwt token to access.