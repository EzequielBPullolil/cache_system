# Cache System

Cache system developed in golang

## Description

The cache system is an application developed in Golang, consisting of two main components. The first is the TCP server, enabling anyone with the host and port to utilize cache system operations. The second part is the interactive shell, facilitating client communication with the TCP server in a user-friendly manner.

### Operations

1. Cache: To cache information explicitly, this operation receives 3 pieces of data
    - uuid: to identify the data
    - data: the data to be cached
    - expiresIn: the time for which the data remains "alive" [comming soon!]
2. Uncache: To uncaché information explicitly, this operation takes a single parameter, which is the uuid identifying the data.
3. Fetch: To retrieve the cached data by the identifier assigned during caching, this operation takes a single parameter, which is the uuid of the data.

## Install
First and foremost, we need to have Go installed in version 1.21.x.
1. Clone repository
    ```
    git clone https://github.com/EzequielBPullolil/cache_system && cd cache_system
    ```
2. Run the install script 
    ```
    source ./install.sh
    ```
This will install the cache system in the directory $HOME/.cache_system_ and add the commands 'cache_server' to execute the server, and 'cachex' for the interactive shell.


## Use cache system
To utilize the cache system, we must first initiate the TCP server. To do so, execute the following command:
```
cache_server start [host] [port]
```
This will start the TCP server of the cache system, allowing anyone with the host and port to utilize the cache system

```
echo "cache hello world" | nc [host] [port] 
```
By default, if the 'host' and 'port' parameters are not provided, the TCP server and shell use the following address: localhost:290333
## Shell

To initiate the shell, you need to provide the host and port of the TCP server.

```
cachex --host=[host] --port=[port] 
```

Shell example
```
Cache shell 
cachex> cache hello world
Cached hello world
cache>
```


---
```
Versions
[Go 1.21.6](https://go.dev/dl/)
```