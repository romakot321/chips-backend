package models

import (
  "encoding/json"
)

func ToMap(strct interface{}) map[string]interface{} {
  var out map[string]interface{}
  temp, _ := json.Marshal(strct)
  json.Unmarshal(temp, &out)
  return out
}
