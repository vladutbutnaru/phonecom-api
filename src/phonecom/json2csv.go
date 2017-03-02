package main

import (
  "encoding/json"
  "encoding/csv"
  "strings"
  "math"
  "fmt"
  "os"
  "flag"
  "unicode/utf8"
)

func json2csv(output []byte, keys []string) {

  var writer *csv.Writer
  writer = csv.NewWriter(os.Stdout)
  outputDelim := flag.String("d", ",", "delimiter used for output values")
  delimiter, _ := utf8.DecodeRuneInString(*outputDelim)
  writer.Comma = delimiter

  var expanded_keys [][]string
  for _, key := range keys {
    expanded_keys = append(expanded_keys, strings.Split(key, "."))
  }

  writer.Write(keys)
  writer.Flush()

  var data map[string]interface{}
  json.Unmarshal(output, &data)

  var record []string
  for _, expanded_key := range expanded_keys {
    record = append(record, get_value(data, expanded_key))
  }

  writer.Write(record)
  writer.Flush()
}

func get_value(data map[string]interface{}, keyparts []string) string {

  if len(keyparts) > 1 {

    subdata, _ := data[keyparts[0]].(map[string]interface{})
    return get_value(subdata, keyparts[1:])

  } else if v, ok := data[keyparts[0]]; ok {

    switch v.(type) {
    case nil:
      return ""
    case float64:
      f, _ := v.(float64)
      if math.Mod(f, 1.0) == 0.0 {
        return fmt.Sprintf("%d", int(f))
      } else {
        return fmt.Sprintf("%f", f)
      }
    default:
      return fmt.Sprintf("%+v", v)
    }
  }

  return ""
}

type StringArray []string

func (a *StringArray) Set(s string) error {
  for _, ss := range strings.Split(s, ",") {
    *a = append(*a, ss)
  }
  return nil
}

func (a *StringArray) String() string {
  return fmt.Sprint(*a)
}