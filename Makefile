PROTO_DIR=greet
GRPC_DIR=greetsvc	# or say `SVC_DIR`

PROTO_SRC=$(wildcard ${PROTO_DIR}/*.proto)
PROTO_GENOPT=--go_opt=paths=source_relative
PROTO_OPT=--go_out=.
PROTO_INC=-I=.
PROTO_PB=$(patsubst %.proto,%.pb.go,${PROTO_SRC})

GRPC_SRC=$(wildcard ${GRPC_DIR}/*.proto)
GRPC_GENOPT=--go-grpc_opt=paths=source_relative
GRPC_OPT=--go-grpc_out=.
GRPC_PB=$(patsubst %.proto,%.pb.go,${GRPC_SRC})
GRPC_RES=$(patsubst %.proto,%_grpc.pb.go,${GRPC_SRC})

CLIENT_SRC=$(wildcard client/*.go)
SERVER_SRC=$(wildcard server/*.go)

.PHONY: proto grpc server client clean tidy

# --- recipes

${PROTO_PB}: ${PROTO_SRC}
	protoc ${PROTO_INC} ${PROTO_OPT} ${PROTO_GENOPT} $^
proto: ${PROTO_PB}

${GRPC_RES}: ${GRPC_SRC} ${PROTO_PB}
	protoc ${PROTO_INC} ${PROTO_OPT} ${GRPC_OPT} \
		${PROTO_GENOPT} ${GRPC_GENOPT} ${GRPC_SRC}
grpc: ${GRPC_RES}

client: ${CLIENT_SRC}
client: ${GRPC_RES}
client:
	go run ${CLIENT_SRC}

server: ${SERVER_SRC}
server: ${GRPC_RES}
server: 
	-go run ${SERVER_SRC}

clean:
	rm -rf ${PROTO_DIR}/*.pb.go ${GRPC_DIR}/*.pb.go

tidy:
	go mod tidy
