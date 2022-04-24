alias fixprisma="mkdir $PWD/prisma/db" # fixes prisma generate unable to create .gitignore
alias fixbashrc="sed -i 's/\r$//' $BASH_SOURCE" # fixes dos2unix \r problem
alias updatesource="source $BASH_SOURCE"
alias prisma="go run github.com/prisma/prisma-client-go"
alias migrate="prisma migrate dev --name init"
alias generate="prisma generate"
alias getswag="go install github.com/swaggo/swag/cmd/swag@latest"