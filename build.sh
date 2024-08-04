
# Build Go version
protoc --go_out=go_version/ --go-grpc_out=go_version/ proto/model.proto 
cd go_version
CGO_ENABLED=0 go build -ldflags "-s -w" -o ../bin/go_server cmd/server/main.go
cd ..


# Build rust
cd rust_version
cargo build --release --bin rust_server
cd ..
mv rust_version/target/release/rust_server bin/rust_server
