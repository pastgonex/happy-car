#! /bin/zsh

function genProto {
  # shellcheck disable=SC2034
  DOMAIN=$1
  SKIP_GATEWAY=$2
  PROTO_PATH=./${DOMAIN}/api
  GO_OUT_PATH=./${DOMAIN}/api/gen/v1
  mkdir -p "${GO_OUT_PATH}"

  protoc -I="${PROTO_PATH}" --go_out=paths=source_relative:"${GO_OUT_PATH}" --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:"${GO_OUT_PATH}" "${DOMAIN}".proto

  # shellcheck disable=SC1072
  # shellcheck disable=SC1020
  # shellcheck disable=SC1073
  if [ "$SKIP_GATEWAY" ]; then
    return
  fi
  protoc -I="${PROTO_PATH}" --grpc-gateway_out=paths=source_relative,grpc_api_configuration="${PROTO_PATH}"/"${DOMAIN}".yaml:"${GO_OUT_PATH}" "${DOMAIN}".proto

  # shellcheck disable=SC2034
  PBTS_BIN_DIR=../wx/miniprogram/node_modules/.bin
  # shellcheck disable=SC2034
  PBTS_OUT_DIR=../wx/miniprogram/service/proto_gen/${DOMAIN}
  mkdir -p "$PBTS_OUT_DIR"
  ${PBTS_BIN_DIR}/pbjs -t static -w es6 "${PROTO_PATH}"/"${DOMAIN}".proto --no-create --no-encode --no-decode --force-number --no-verify --no-delimited > "${PBTS_OUT_DIR}"/"${DOMAIN}"_pb_temp.js
  # shellcheck disable=SC2016
  # shellcheck disable=SC2028
  # shellcheck disable=SC2086
  echo 'import * as $protobuf from "protobufjs";\n' > ${PBTS_OUT_DIR}/"${DOMAIN}"_pb.js
  cat "${PBTS_OUT_DIR}"/"${DOMAIN}"_pb_temp.js >> "${PBTS_OUT_DIR}"/"${DOMAIN}"_pb.js
  # shellcheck disable=SC2086
  rm ${PBTS_OUT_DIR}/${DOMAIN}_pb_temp.js
  # shellcheck disable=SC2086
  ${PBTS_BIN_DIR}/pbts -o ${PBTS_OUT_DIR}/"${DOMAIN}"_pb.d.ts ${PBTS_OUT_DIR}/"${DOMAIN}"_pb.js
}

genProto auth
genProto rental
genProto blob 1
genProto car