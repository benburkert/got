#!/usr/bin/env bats

load test_helper

repo "go"

@test "got log" {
  assert log
}

@test "got log --abbrev-commit" {
  assert log --abbrev-commit
}

@test "got log --pretty=oneline" {
  assert log --pretty=oneline
}

@test "got log --format=oneline" {
  assert log --format=oneline
}
