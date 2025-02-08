#!/bin/bash
###
 # @Author: javohir-a abdusamatovjavohir@gmail.com
 # @Date: 2024-11-22 23:17:53
 # @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 # @LastEditTime: 2024-11-22 23:23:31
 # @FilePath: /admin/pkg/scripts/gen_proto.sh
 # @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
### 
echo "Running script with argument: $1"
CURRENT_DIR=$1  
GEN_OUT="${CURRENT_DIR}/genproto"

# Clean the genproto directory before generating
rm -rf "${GEN_OUT}"
mkdir -p "${GEN_OUT}"

# Find and compile all .proto files
find "${CURRENT_DIR}/protos" -type f -name "*.proto" -print0 | while IFS= read -r -d '' file; do
  protoc \
    -I="${CURRENT_DIR}/protos" \
    --go_out="${GEN_OUT}" \
    --go-grpc_out="${GEN_OUT}" \
    "${file}"
done
