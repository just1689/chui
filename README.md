# chisel-tui
A small wrapper about chisel. The goal of this project is to make it easier to work with chisel for development purposes for example.

## Generate config file
```bash
    # Generate the file from the executable
    chui -g config.json
    # or if you have the project
    go run app.go -g config.json
```

## Run locally
```bash
    # Run the file from the executable
    chui -c config.json
    # or if you have the project
    go run app.go -c config.json
```

## Run in Docker
```bash
    # Make sure you have generated the file
    # chui -g config.json

    # Build the declarative container
    docker build -t t:t --file Dockerfile.declarative .
```



