# Dumb Redis

Dummy project which accepts connection from redis, relied on [this tutorial](https://www.build-redis-from-scratch.dev/en/resp-reader)


## Getting Started

### Building
#### With Makefile
Simply run the following commands:

```bash
$ make
```

#### With go cli

```bash
$ go build
```

### Running
After building, the executable output file would exist in `./bin`, simply run that file

## Connecting to Redis

### Windows
1. Ensure you have [WSL installed](ttps://learn.microsoft.com/en-us/windows/wsl/install)
2. Run the executable file in the `./bin/` directory
3. Install any Linux distribution (e.g. Ubuntu)
4. Install [Redis CLI](https://redis.io/docs/latest/operate/oss_and_stack/install/archive/install-redis/install-redis-on-linux/) on the machine
5. Run the following command to get the ip address of your windows machine 
```bash
ip route show | grep -i default | awk '{ print $3}'
``` 
A sample output would look like this
```
172.30.96.1
```

6. Run the following command to connect to your host's server:
```bash
redis-cli -u redis://172.30.96.1:6379
```








