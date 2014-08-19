GopherBlog
==========

My blog about Go, written in Go. Located at www.gopherblog.org

### Requirements

The system assumes that you have successfully installed the following requirements:

* A running version of [MySQL](http://dev.mysql.com/doc/refman/5.1/en/installing.html)
* The [Go](https://golang.org/doc/install) language set-up, with correct environment variables (GOPATH, GOROOT, etc)

### Setting up the database

To set up the database you need to run the following command:

`chmod +x install.sh; ./install.sh`

Follow the on-screen instructions and enter your database host, and credentials. The second section will ask you to create an admin username which you will use to create posts on the blog.

### Running the blog

First we must compile the main package and all its dependencies into a binary. To do this we run:

`go build`

This should create the binary called `gopherblog`. On my machine, for example, I naively run the blog as:

`./gopherblog -db="root:root@tcp(localhost:3306)/gopherblog"`

The following command-line flags are provided:

| Flag         | Default Value                              | Description                                     |
| ------------- |:------------------------------------------|:------------------------------------------------|
| `db`           | user:pass@tcp(localhost:3306)/gopherblog | Database connection string.   |
| `host`           | localhost | Hostname that the server runs on |
| `port`           | 8080 | Port to listen on for HTTP connections |
| `nocache`        |  | Setting this flag will recompile the HTML templates on every request |


### Navigating

| Route         | Description                                     |
| ------------- |:------------------------------------------------|
| `/`           | Home page of the blog, displays a list of posts |
| `/post/{slug}`| Displays a post with the given slug             |
| `/login`      | Displays login page                             |
| `/manage`     | Manage your posts (requires authentication)     |
| `/edit/{id}`  | Edit a post with the given ID                   |
