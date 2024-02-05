# Cache System

Cache system developed in golang

## Description

The cache system is an application developed in Golang, consisting of two main components. The first is the TCP server, enabling anyone with the host and port to utilize cache system operations. The second part is the interactive shell, facilitating client communication with the TCP server in a user-friendly manner.

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

## Operations
Operations can be performed using the shell or by making requests directly to the TCP server.

Here is an example with the 3 available operations of the system, both using the shell and making a request


### cache 

The cache operation requires two arguments, the UUID, and the value to which the UUID refers.

In the following example, I use 'hello' as the UUID and 'world' as the value to associate.

#### With simple request
``` bash
echo "cache hello world" | nc [host] [port] 
```
#### Shell example
```sql
Cache shell 
cachex> cache hello world
```
---
### fetch 
The fetch operation is used to retrieve the data associated with a UUID. This operation only requires a single argument, which is the UUID of the data. If there is no data associated with the UUID, it does nothing.

In the following example, I use 'hello' as the UUID and 'world' as the value to associate.

#### With simple request
``` bash
echo "fetch hello " | nc [host] [port] 
world
```
#### Shell example
```
Cache shell 
cachex>fetch hello 
hello=world
```
---
### uncache 
The uncache operation requires a single argument, which is the UUID of the data to uncaché. If there is no cached data with the given UUID, it does nothing.

In the following example, I use 'hello' as the UUID

#### With simple request
``` bash
echo "uncache hello" | nc [host] [port] 
```
#### Shell example
```sql
Cache shell 
cachex> uncache hello
```

---
```
Versions
[Go 1.21.6](https://go.dev/dl/)
```