package prisma

import (
	"boilerplate/prisma/db"
	"context"
)

var (
	Ctx    = context.Background()
	Client = db.NewClient()
	_      = Client.Prisma.Connect()
)
