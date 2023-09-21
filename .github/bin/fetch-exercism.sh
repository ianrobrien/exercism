#!/bin/bash

# Define the URL, version, and expected SHA256 hash as variables
base_url="https://github.com/exercism/cli/releases/download"
version="3.2.0"
filename="exercism-${version}-linux-x86_64.tar.gz"
url="${base_url}/v${version}/${filename}"

# Expected SHA256 hash (replace with the actual expected hash)
expected_hash="YOUR_EXPECTED_SHA256_HASH_HERE"

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
tar -xzvf "$filename"
chmod +x "$filename"
