#!/usr/bin/env bash
# Script to build and install gocapsule with the correct Go version
# This is a workaround for gocapsule version compatibility issues

set -e

GOCAPSULE_VERSION="v0.1.2"
TEMP_DIR=$(mktemp -d)
GOBIN=${GOBIN:-${GOPATH:-$HOME/go}/bin}

echo "Building gocapsule from source..."

# Clone the repository
git clone --depth 1 --branch ${GOCAPSULE_VERSION} https://github.com/YuitoSato/gocapsule.git "$TEMP_DIR/gocapsule" 2>/dev/null || \
  git clone --depth 1 https://github.com/YuitoSato/gocapsule.git "$TEMP_DIR/gocapsule"

cd "$TEMP_DIR/gocapsule"

# Update go.mod to use Go 1.25.4 (matching the project's Go version)
go mod edit -go=1.25.4
go mod tidy

# Build and install
go build -o "$GOBIN/gocapsule" .

# Cleanup
rm -rf "$TEMP_DIR"

echo "✓ gocapsule installed to $GOBIN/gocapsule"
