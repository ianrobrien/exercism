#!/bin/bash

# Initialize the exit code to 0
exit_code=0

# Iterate through all subdirectories of the "go" directory
for exercise in go/*/; do
  # Change to the exercise directory
  pushd "$exercise" || exit 1

  # Run golangci-lint quietly and capture its exit code
  if golangci-lint run ./...; then
    : # Do nothing on success
  else
    # If it fails, print the output and set exit_code to 1
    echo "Linting failed in $(pwd):"
    golangci-lint run ./...
    exit_code=1
  fi

  # Return to the original directory
  popd || exit 1
done

# Return the exit code at the end of the script
exit $exit_code
