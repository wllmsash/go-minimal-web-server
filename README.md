# go-minimal-web-server

Run the web server in development:

```sh
go run .
```

Build and run an ephemeral web server:

```sh
docker build \
  --tag go-minimal-web-server:latest \
  .

docker run \
  --rm \
  --publish 8080:8080 \
  --mount type=bind,source=/var/local/www,target=/var/local/www,readonly \
  go-minimal-web-server:latest
```

## Developing in a devcontainer

Build the devcontainer image:

```sh
docker build \
  --file ./.devcontainer/Dockerfile \
  --tag go-minimal-web-server-devcontainer:latest \
  --build-arg DEVCONTAINER_UID=$UID \
  .
```

Run the devcontainer from VSCode, or start one from the terminal:

```sh
docker run \
  --rm \
  --tty \
  --interactive \
  --publish 8080:8080 \
  --mount type=bind,source=$PWD,target=/workspace \
  --add-host=host.docker.internal:host-gateway \
  go-minimal-web-server-devcontainer:latest \
  bash
```
