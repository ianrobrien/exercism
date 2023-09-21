#!/bin/bash

# Define the URL, version, and expected SHA256 hash as variables
base_url="https://github.com/golangci/golangci-lint/releases/download/"
version="1.54.2"
filename="golangci-lint-${version}-linux-amd64.tar.gz"
url="${base_url}/v${version}/${filename}"

# Expected SHA256 hash (replace with the actual expected hash)
expected_hash="17c9ca05253efe833d47f38caf670aad2202b5e6515879a99873fabd4c7452b3"

# Use curl to download the file
curl -sSfL -o "$filename" "$url"

# Compute the SHA256 hash of the downloaded file
computed_hash=$(sha256sum "$filename" | awk '{print $1}')

# Compare the computed hash with the expected hash
if [ "$computed_hash" == "$expected_hash" ]; then
    echo "File integrity verified. Hash matches: $computed_hash"
else
    echo "File integrity check failed. Hash does not match."
    exit 1
fi

# Extract the tar.gz file to the current directory
mkdir -p bin && tar -xzvf "$filename"
mv golangci-lint-${version}-linux-amd64/golangci-lint bin/

