Build server:
```
cd server
docker build --progress=plain --no-cache . -t my_server
```

# Build client:
```
cd client
docker build --progress=plain --no-cache . -t my_client
```

# How to run this code

We need to have the two docker containers run on the same network. 

```bash
docker network create --subnet=172.18.0.0/16 my_network
```


## Run server on a static IP on this network
```
docker run --rm --net my_network --ip 172.18.0.2 --name my_server my_server

```

## Run client on the same network
```
docker run --rm --net my_network -p 5300:5300 --name my_client my_client
```
