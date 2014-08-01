package blog

// Holds user (or author) data
type User struct {
	id                    int
	name, email, password string
	isAuthor              bool
}
