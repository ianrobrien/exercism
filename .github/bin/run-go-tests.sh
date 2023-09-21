#!/bin/bash

exercism_version="3.2.0"
exercism_expected_hash="4ea3e1ea8916a8003da95dbd6eef7a3a29802e637ed6a0f2aaaa2f1c98754915"
exercism_base_url="https://github.com/exercism/cli/releases/download/"
exercism_filename="exercism-${exercism_version}-linux-x86_64.tar.gz"
exercism_url="${exercism_base_url}/v${exercism_version}/${exercism_filename}"

golangci_lint_version="1.54.2"
golangci_lint_expected_hash="17c9ca05253efe833d47f38caf670aad2202b5e6515879a99873fabd4c7452b3"
golangci_lint_base_url="https://github.com/golangci/golangci-lint/releases/download/"
golangci_lint_filename="golangci-lint-${golangci_lint_version}-linux-amd64.tar.gz"
golangci_lint_url="${golangci_lint_base_url}/v${golangci_lint_version}/${golangci_lint_filename}"

# Initialize the exit code to 0
exit_code=0

# Function to download, verify, and extract a tar.gz file
download_verify_extract() {
    local url="$1"
    local filename="$2"
    local expected_hash="$3"

    # Use curl to download the file
    curl -sSfL -o "$filename" "$url"

    # Compute the SHA256 hash of the downloaded file
    local computed_hash
    computed_hash=$(sha256sum "$filename" | awk '{print $1}')

    # Compare the computed hash with the expected hash
    if [ "$computed_hash" == "$expected_hash" ]; then
        echo "File integrity verified. Hash matches: $computed_hash"
    else
        echo "File integrity check failed. Hash does not match."
        exit 1
    fi

    # Extract the tar.gz file to the current directory
    mkdir -p bin
    if tar -tzf "$filename" | grep -q '/'; then
        tar -xzf "$filename" -C bin --strip-components=1
    else
        tar -xzf "$filename" -C bin/
    fi

    export PATH="$GITHUB_WORKSPACE/bin:$PATH"
}

# Function to lint Go source files in a directory
lint_go_files() {
  for exercise in go/*/; do
    pushd "$exercise" || exit 1

    if ! golangci-lint run .; then
        echo "Linting for $PWD failed."
        exit_code=1
    fi

    # Return to the original directory
    popd || exit 1
  done
}

# Function to run "exercism test" in a directory
run_exercism_test() {
  for exercise in go/*/; do
    pushd "$exercise" || exit 1

    # Run "exercism test" and capture its exit code
    if exercism test; then
      : # Do nothing on success
    else
      # If it fails, print the output and set exit_code to 1
      echo "Testing failed in $(pwd):"
      exit_code=1
    fi

    # Return to the original directory
    popd || exit 1
  done
}

download_verify_extract "$exercism_url" "$exercism_filename" "$exercism_expected_hash"
download_verify_extract "$golangci_lint_url" "$golangci_lint_filename" "$golangci_lint_expected_hash"

# Iterate through all subdirectories of the "go" directory
echo "linting all go files"
lint_go_files "$exercise"
echo "running exercism tests"
run_exercism_test "$exercise"

# Return the exit code at the end of the script
exit $exit_code
