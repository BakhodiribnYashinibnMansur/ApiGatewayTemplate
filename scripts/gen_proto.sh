# !/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
done


# sed -i "" -e "s///g" ./api/proto/*.go
