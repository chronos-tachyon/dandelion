#!/bin/bash
set -euo pipefail
umask 022
cd "$(dirname "$0")"
find . -name "*.pb.go" -type f -delete
find proto -name "*.proto" -type f -print0 | sed -ze 's|^proto/||' | xargs -0 protoc \
  -Iproto \
  --go_out=module=github.com/chronos-tachyon/dandelion:. \
  --go-vsrpc_out=module=github.com/chronos-tachyon/dandelion:.
