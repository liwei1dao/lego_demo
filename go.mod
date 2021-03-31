module lego_demo

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/liwei1dao/lego v0.0.0-20200513072210-02b6de94b6ca
	go.mongodb.org/mongo-driver v1.4.6
)

replace github.com/liwei1dao/lego => ../lego
