# chui
A small wrapper around <a href="https://github.com/jpillora/chisel">Chisel</a>. The goal of this project is to make it easier to work with chisel in a declarative way.

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

    # ... and edited also
    # vi config.json

    # Build the declarative container
    docker build -t image:tag --file Dockerfile.declarative .
    docker run -p port:port image:tag
```

## Usage examples
- Proxy a number of services on another network - in a container cluster for example.
- Expose a service in one cluster to another cluster.
- Development

