# Game Engine

## API

You can check out the API specification in http://editor.swagger.io/ and importing the following URL:

```
https://raw.githubusercontent.com/VxRackNeutrino/Hero/develop/game-engine/swagger.yaml
```

## Development

For development purposes, we recommend using the Vagrant box available on the project root directory. That will install and set up Golang, Docker and a MySQL database needed to work on the Game Engine development.

Once you have the Vagrant box set up, ssh into it and CD to the `game-engine` directory.

Then, export the database configuration by exporting the following variable.

```
export DATABASE_URL=root:root@/titandb
```

And, optionally, the admin token used to issue API requests that require admin privileges.

```
export ADMIN_TOKEN=1234
```

Install the Golang dependencies.

```
go get -v ./...
```

Then, you can run the Game Engine with the following command.

```
go run *.go
```

## Docker

You can also build a Docker image of the Game Engine.

```
docker build -t game-engine .
```

And run it.

```
docker run -d \
  --name game-engine \
  -e DATABASE_URL=root:root@/titandb \
  -p 0.0.0.0:80:8080 \
   game-engine
```

You can see more examples of database URLs here: https://github.com/go-sql-driver/mysql#dsn-data-source-name
