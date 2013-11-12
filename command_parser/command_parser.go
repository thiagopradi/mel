package command_parser;

import(
  "strings"
)

type ServerStorage struct {
  Dict map[string]string 
}

func (s *ServerStorage) ParseCommand(command string) (response string) {
  if s.Dict == nil {
    s.Dict = make(map[string]string)
  }

  arr := strings.Split(command, " ")

  if arr[0] == "SET" {
    arr := strings.Split(command, " ")
    s.Dict[arr[1]] = arr[2]
    return "OK"
  } else if arr[0] == "GET" {
    return s.Dict[arr[1]]
  } else if arr[0] == "DEL" {
    delete(s.Dict, arr[1])
    return "OK"
  }

  return "command failed"
}
