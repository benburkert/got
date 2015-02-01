#!/usr/bin/env bats

load test_helper

repo "go"

@test "got log" {
  assert log
}

@test "got log --abbrev-commit" {
  assert log --abbrev-commit
}
