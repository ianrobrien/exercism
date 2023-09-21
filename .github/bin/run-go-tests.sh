#!/bin/bash

# Initialize the exit code to 0
exit_code=0

# Function to lint Go source files in a directory
lint_go_files() {
  local dir="$1"
  pushd "$dir" || exit 1

  # Use find to filter Go source code files for linting (exclude test files)
  files_to_test=$(find . -type f -name "*.go" ! -name "*_test.go")

  # Run golangci-lint quietly and capture its exit code
  if [ -n "$files_to_test" ]; then
    if golangci-lint run $files_to_test; then
      : # Do nothing on success
    else
      # If it fails, print the output and set exit_code to 1
      echo "Linting failed in $(pwd):"
      golangci-lint run $files_to_test
      exit_code=1
    fi
  fi

  # Return to the original directory
  popd || exit 1
}

# Function to run "exercism test" in a directory
run_exercism_test() {
  echo $PATH
  local dir="$1"
  pushd "$dir" || exit 1

  # Run "exercism test" and capture its exit code
  if exercism test; then
    : # Do nothing on success
  else
    # If it fails, print the output and set exit_code to 1
    echo "Testing failed in $(pwd):"
    exercism test
    exit_code=1
  fi

  # Return to the original directory
  popd || exit 1
}

# Iterate through all subdirectories of the "go" directory
for exercise in go/*/; do
#  lint_go_files "$exercise"
  run_exercism_test "$exercise"
done

# Return the exit code at the end of the script
exit $exit_code
