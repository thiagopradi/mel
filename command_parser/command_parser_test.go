package command_parser;

import (
  "testing"
)

func Test_parseSetCommand (t *testing.T) {
  ss :=  new(ServerStorage)

  str := ss.ParseCommand("SET thiagopradi valor")

  if ss.Dict["thiagopradi"] != "valor" {
    t.Error("failed to set thiagopradi key")
  }

  if str != "OK" {
    t.Error("failed to get return value")
  }

  return_value := ss.ParseCommand("GET thiagopradi")

  if return_value != "valor" {
    t.Error("Failed to get key thiagopradi")
  }
}
