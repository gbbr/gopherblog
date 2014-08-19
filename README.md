GopherBlog
==========

My blog about Go, written in Go.  
Located at www.gopherblog.org

### Requirements

The system assumes that you have successfully installed the following requirements:

* [MySQL](http://dev.mysql.com/doc/refman/5.1/en/installing.html)
* [Go](https://golang.org/doc/install)

### Setting up the database

To set up the database you need to run the following command:

`chmod +x install.sh; ./install.sh`

Follow the on-screen instructions and enter your database host, and credentials. The second section will ask you to create an admin username which you will use to create posts on the blog.

### Running the blog

First we must compile the main package and all its dependencies into a binary. To do this we run:

`go build`

This should create the binary called `gopherblog` for you which comes with the following flags:

| Flag         | Default Value                              | Description                                     |
| ------------- |:------------------------------------------|:------------------------------------------------|
| `db`           | user:pass@tcp(localhost:3306)/gopherblog | Database connection information. Make sure you enter the same string, but containing your own username and password  |
| `host`           | localhost | Hostname that the server runs on |
| `port`           | 8080 | Port to listen on for HTTP connections |
| `nocache`        | `false` | Setting this flag will recompile the HTML templates on every request |

On my machine, for example, I run the blog as:

`./gopherblog -db=root:root@tcp(localhost:3306)/gopherblog`

My username and password is cleverly named `root` and `root`

### Navigating

| Route         | Description                                     |
| ------------- |:------------------------------------------------|
| `/`           | Home page of the blog, displays a list of posts |
| `/post/{slug}`| Displays a post with the given slug             |
| `/login`      | Displays login page                             |
| `/manage`     | Manage your posts (requires authentication)     |
| `/edit/{id}`  | Edit a post with the given ID                   |
