# grpc-piglow
Connect to your raspberry Pi PiGlow using [gRPC](http://www.grpc.io)

## Goal
This library will help making RPC calls to a remote Pi with a PiGlow on your network.

## Client and Server implementation
You can find 2 reference implementations for a client and server in the *examples/* directory. Those 2 are using the *bonjour*
protocol to discover each other, and then, once connected, you can do direct calls to the remote Pi which will be reflected
with low latency.

The API for accessing PiGlow is mirroring the excellent [go-piglow](https://github.com/wjessop/go-piglow) API, adding error
and out of bounds handling. It then uses that library to deal with the PiGlow.

## Other languages
As we are using gRPC and protobuf, you can write a client in any languages you need. Just use the [piglow.proto file](proto/piglow.proto)
to generate your client code.

## Snap
If you are using [Ubuntu Core](https://developer.ubuntu.com/en/snappy/start/raspberry-pi-2/) on your Pi, you can directly
install the server snap with
`snap install grpc-piglow`
