#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Set WASM target
export GOOS=js
export GOARCH=wasm

# Build the WASM file
echo "Building WebAssembly file..."
go build -o main.wasm main.go

# Copy wasm_exec.js from Go installation
#cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

echo "Build complete! Files created:"
echo "  - main.wasm (WebAssembly binary)"
echo "  - wasm_exec.js (Go WASM runtime)"
echo "  - index.html (Web page)"

# Instructions
echo ""
echo "To run:"
echo "1. Start a local web server:"
echo "   python3 -m http.server 8080"
echo "2. Open http://localhost:8080 in your browser"
