#!/usr/bin/env bash

main() {
  string1="$1"
  string2="$2"
  distance=0

  # Check if the number of arguments is correct
  [ "$#" -ne 2 ] && echo "Usage: hamming.sh <string1> <string2>" && exit 1

  # Check if the strings are of equal length
  [ "${#string1}" -ne "${#string2}" ] && echo "strands must be of equal length" && exit 1

  # Calculate hamming distance
  for ((i = 0; i < ${#string1}; i++)); do
    char1="${string1:i:1}"
    char2="${string2:i:1}"

    if [ "$char1" != "$char2" ]; then
      ((distance++))
    fi
  done

  echo "$distance"
}

main "$@"
