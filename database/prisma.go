package prisma

import (
	"boilerplate/prisma/db"
	"context"
)

var (
	Client = db.NewClient()
	_      = Client.Prisma.Connect()
	Ctx    = context.Background()
)
