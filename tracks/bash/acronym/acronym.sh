#!/usr/bin/env bash

main() {
  string="$1"
  acronym=""
  IFS="- " read -ra arr <<<"$string"

  for i in "${arr[@]}"; do
    acronym+="${i:0:1}"
  done

  echo "$acronym"
}

main "$@"
