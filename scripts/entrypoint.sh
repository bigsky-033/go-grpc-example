#!/bin/sh

GRPC_PORT=${GRPC_PORT:-18080}
HTTP_HEALTH_CHECK_PORT=${HTTP_HEALTH_CHECK_PORT:-8081}

/app/program/server \
  --grpc_port="${GRPC_PORT}" \
  --http_health_check_port="${HTTP_HEALTH_CHECK_PORT}"
