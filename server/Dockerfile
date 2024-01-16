# First stage: Build the Protobuf files using Buf
FROM bufbuild/buf:latest as bufbuild
WORKDIR /app
COPY . .
RUN buf generate proto

# Second stage: Build the Go binary
FROM golang:1.21 as gobuild
WORKDIR /app
# Copy the generated files from the previous stage
COPY --from=bufbuild /app .
# Copy the rest of the Go source code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server server/main.go

# Final stage: Create the runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the compiled binary from the second stage
COPY --from=gobuild /app/server .
EXPOSE 50051
CMD ["./grpc-server"]
