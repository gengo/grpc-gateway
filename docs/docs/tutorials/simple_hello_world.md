---
layout: default
title: Creating a simple hello world with gRPC
parent: Tutorials
nav_order: 2
---

### Creating a simple hello world with gRPC

### Define your gRPC service using protocol buffers

Before we create a gRPC service, we should create a proto file to define what we need, here we create a file named `hello_world.proto` in the directory `proto/helloworld/hello_world.proto`.

```proto
syntax = "proto3";

package helloworld;

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```