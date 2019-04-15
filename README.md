# google-oss-stack
An experimental repo building app using only Google open source tool stack.

This is a sandbox for trying out Google open source tools, and integrating the various tools in a single project, from message serialization/deserialization, to communication (RPC), to creating reproducible build, and finally, to packaging and release.

The project will be wrapped inside a Docker container for isolation purpose.

## Tools List
- Protocol Buffers
- Bazel
- gRPC

## Dev Guide
To start a new Docker container with everything set up, run `make start`
Log into the container, and play with the various tools, such as `bazel test //...` or `protoc ...`

To run the test, run `make test`

## References
This project references the various documentation and tutorials from `developers.google.com`
The demo project comes from the Protocol Buffer's Go tutorial.
- https://developers.google.com/protocol-buffers/docs/gotutorial
