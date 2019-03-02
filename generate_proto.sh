#!/bin/bash
echo "Generating proto grpc files..."
protoc ./api/api.proto --proto_path=. --go_out=plugins=grpc:./
echo "done"