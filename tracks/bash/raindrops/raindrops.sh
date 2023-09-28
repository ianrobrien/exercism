#!/usr/bin/env bash

main() {
  result=""
  number=$1

  if [ "$((number % 3))" -eq 0 ]; then
    result+="Pling"
  fi
  if [ "$((number % 5))" -eq 0 ]; then
    result+="Plang"
  fi
  if [ "$((number % 7))" -eq 0 ]; then
    result+="Plong"
  fi

  result="$([ -z "$result" ] && echo "$number" || echo "$result")"
  echo "$result"
}

main "$@"
