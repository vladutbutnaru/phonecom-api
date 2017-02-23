package main

import (
  "testing"
)

func TestCreateGetReplaceDevice(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCliWithJsonIn(createDevice, "../test/jsonin/createDevice.json")
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}
