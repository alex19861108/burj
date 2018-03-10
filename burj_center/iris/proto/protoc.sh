protoc --go_out=plugins=grpc:. $1.proto
protoc-go-inject-tag -input=./$1.pb.go