#!/bin/bash

files='(/cmd|/doc|/infra|/swagger|/src/external/database|/src/external/handler|/http_server|api|/src/pkg/uuid/mock|/src/external/handler/http_server|/src/application/contract/mock|/src/external/database/postgres|/src/external/handler/sqs)'
go list ./... | egrep -v $files$\
