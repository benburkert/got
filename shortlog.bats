#!/usr/bin/env bats

load test_helper

repo "go"

@test "got shortlog" {
  assert shortlog
}
