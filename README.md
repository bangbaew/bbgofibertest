# GoFiber Docker Boilerplate

![Release](https://img.shields.io/github/release/gofiber/boilerplate.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/boilerplate/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/boilerplate/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/boilerplate/workflows/Linter/badge.svg)


## IDE Development

### Visual Studio Code

Use the following plugins, in this boilerplate project:
- Nome: Go
  - ID: golang.go
  - Descrição: Rich Go language support for Visual Studio Code
  - Versão: 0.29.0
  - Editor: Go Team at Google
  - Link do Marketplace do VS: https://marketplace.visualstudio.com/items?itemName=golang.Go

## Development

### Start the application 


```bash
go run app.go
```

### Use local container

```
# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requeriments

# Generate docker image
make build

# Generate docker image with no cache
make build-no-cache

# Run the projec in a local container
make up

# Run local container in background
make up-silent

# Stop container
make stop

# Start container
make start
```

## Production

```bash
docker build -t gofiber .
docker run -d -p 3000:3000 gofiber
```

Go to http://localhost:3000:


![Go Fiber Docker Boilerplate](./go_fiber_boilerplate.gif)
