#! /bin/zsh
PROTO_PATH=./auth/api
GO_OUT_PATH=./auth/api/gen/v1
mkdir -p ${GO_OUT_PATH}

protoc -I=${PROTO_PATH} --go_out=paths=source_relative:${GO_OUT_PATH} --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${GO_OUT_PATH} auth.proto
protoc -I=${PROTO_PATH} --grpc-gateway_out=paths=source_relative,grpc_api_configuration=${PROTO_PATH}/auth.yaml:${GO_OUT_PATH} auth.proto

# shellcheck disable=SC2034
PBTS_BIN_DIR=../wx/miniprogram/node_modules/.bin
# shellcheck disable=SC2034
PBTS_OUT_DIR=../wx/miniprogram/service/proto_gen/auth
mkdir -p $PBTS_OUT_DIR
${PBTS_BIN_DIR}/pbjs -t static -w es6 ${PROTO_PATH}/auth.proto --no-create --no-encode --no-decode --no-verify --no-delimited > ${PBTS_OUT_DIR}/auth_pb_temp.js
# shellcheck disable=SC2016
# shellcheck disable=SC2028
echo 'import * as $protobuf from "protobufjs";\n' > ${PBTS_OUT_DIR}/auth_pb.js
cat ${PBTS_OUT_DIR}/auth_pb_temp.js >> ${PBTS_OUT_DIR}/auth_pb.js
rm ${PBTS_OUT_DIR}/auth_pb_temp.js

${PBTS_BIN_DIR}/pbts -o ${PBTS_OUT_DIR}/auth_pb.d.ts ${PBTS_OUT_DIR}/auth_pb.js