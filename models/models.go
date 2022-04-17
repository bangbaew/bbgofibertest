package models

type User struct {
	Firstname string
	Lastname  string
	PicFormat *string
	Address   string
	Tel       string
	Balance   int
	Bio       *string
	Email     string
	Password  string
}

type Artisan struct {
	Firstname string
	Lastname  string
	PicFormat *string
	Tel       *string
	ClusterID string
}

type Cluster struct {
	Name   string
	Region string
	Grade  string
}
