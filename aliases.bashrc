alias fixprisma="mkdir $PWD/prisma/db && touch $PWD/prisma/db/.gitignore"
alias migrate="go run github.com/prisma/prisma-client-go migrate dev --name init"
alias generate="go run github.com/prisma/prisma-client-go generate"
alias getswag="go install github.com/swaggo/swag/cmd/swag@latest"