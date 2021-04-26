# project

two apis that communicate by gRPC. client API also has a rest API.

## docker-compose

Despite having a docker-compose you'll need to build the containers and then run the compose file. After building the services you can run it with:

```shell
    $ docker-compose up
```


## client

    client looks for a file named `ports.json` to parse in its directory, if not present it'll not feed the backend service with data.

### endpoints

> / healtcheck
    should return always 200
> /ports 
    should return 200/500

### build container

```shell
    $ cd client
    $ docker build -t client .
```

## backend

### RPCs

> UpsertPort
    Should try to update or insert a port record

> ListPorts
    Should list all ports in db

### build container

```shell
    $ cd backend
    $ docker build -t backend .
```

