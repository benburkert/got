#!/bin/bash

function repo() {
  pushd "$BATS_TEST_DIRNAME/testdata/$1" 1>/dev/null
}

function assert() {
  run diff -u <(got $@) <(git $@)

  if [ -n "$output" ]; then
    echo "$output" | cdiff -q 2>/dev/null
    return 1
  fi
}
