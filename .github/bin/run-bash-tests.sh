#!/bin/bash

exercism_version="3.2.0"
exercism_expected_hash="4ea3e1ea8916a8003da95dbd6eef7a3a29802e637ed6a0f2aaaa2f1c98754915"
exercism_base_url="https://github.com/exercism/cli/releases/download/"
exercism_filename="exercism-${exercism_version}-linux-x86_64.tar.gz"
exercism_url="${exercism_base_url}/v${exercism_version}/${exercism_filename}"

bats_version="1.10.0"
bats_url="https://github.com/bats-core/bats-core/archive/refs/tags/v${bats_version}.tar.gz"
# Initialize the exit code to 0
exit_code=0

download_verify_extract() {
  local url="$1"
  local filename="$2"
  local expected_hash="$3"

  curl -sSfL -o "$filename" "$url"

  local computed_hash
  computed_hash=$(sha256sum "$filename" | awk '{print $1}')

  if [ "$computed_hash" == "$expected_hash" ]; then
    echo "File integrity verified. Hash matches: $computed_hash"
  else
    echo "File integrity check failed. Hash does not match."
    exit 1
  fi

  mkdir -p bin
  if tar --list -zf "$filename" | grep -q '/' > /dev/null; then
    tar -xzf "$filename" -C bin --strip-components=1
  else
    tar -xzf "$filename" -C bin
  fi

  export PATH="$GITHUB_WORKSPACE/bin:$PATH"
}

download_extract_bats() {
  local url=$bats_url
  local filename="v${bats_version}.tar.gz"

  curl -sSfL -o "$filename" "$url"

  mkdir -p bin
  tar -xzf "$filename" -C
  mv "bats-core-${bats_version}/bin/bats" bin
}

run_exercism_test() {
  for exercise in go/*/; do
    pushd "$exercise" || exit 1

    if ! exercism test; then
      declare -g exit_code=1
    fi

    popd || exit 1
  done
}

download_extract_bats
download_verify_extract "$exercism_url" "$exercism_filename" "$exercism_expected_hash"

run_exercism_test "$exercise"

exit $exit_code
