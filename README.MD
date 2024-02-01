# Cache System

Cache system developed in golang

## Description

The cache system is a TCP/IP server developed in Golang, which consists of the following operations interface:

1. Cache: To cache information explicitly, this operation receives 3 pieces of data
    - uuid: to identify the data
    - data: the data to be cached
    - expiresIn: the time for which the data remains "alive"
2. Uncache: To uncaché information explicitly, this operation takes a single parameter, which is the uuid identifying the data.
3. Fetch: To retrieve the cached data by the identifier assigned during caching, this operation takes a single parameter, which is the uuid of the data.

## Use
 COMMING

## Versions
[Go 1.21.6](https://go.dev/dl/)
