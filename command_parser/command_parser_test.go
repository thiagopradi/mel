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

func Test_parseDelCommand (t *testing.T) {
  ss :=  new(ServerStorage)

  str := ss.ParseCommand("SET thiagopradi valor")

  if ss.Dict["thiagopradi"] != "valor" {
    t.Error("failed to set thiagopradi key")
  }

  if str != "OK" {
    t.Error("failed to get return value")
  }

  new_return := ss.ParseCommand("DEL thiagopradi")

  if new_return != "OK" {
    t.Error("failed to get return value from DEL function")
  }

  new_str := ss.ParseCommand("GET thiagopradi")

  if new_str != "" {
    t.Error("failed to get deleted value")
  }
}
