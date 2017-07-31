#!/bin/bash
echo "Compiling..."
CGO_ENABLED=0 go build && echo "Building..." && docker build -t "steampipe" .
