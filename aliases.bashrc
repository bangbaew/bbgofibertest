export SOURCEFILE=aliases.bashrc
alias fixprisma="mkdir $PWD/prisma/db" # fixes prisma generate unable to create .gitignore
alias updatesource="sed -i 's/\r$//' $SOURCEFILE && source $SOURCEFILE" # fixes win32 invalid char '\r' problem [dos2unix]
alias prisma="go run github.com/prisma/prisma-client-go"
alias migrate="prisma migrate dev --name init"
alias generate="prisma generate"
alias getswag="go install github.com/swaggo/swag/cmd/swag@latest"
echo Loaded $BASH_SOURCE