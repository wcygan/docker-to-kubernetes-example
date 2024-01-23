#!/bin/bash

grpcurl -plaintext -proto proto/ping/v1/ping.proto localhost:8080 ping.v1.PingService/Ping