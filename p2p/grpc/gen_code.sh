#!/usr/bin/env bash
 protoc --go_out=plugins=grpc:. route_guide.proto