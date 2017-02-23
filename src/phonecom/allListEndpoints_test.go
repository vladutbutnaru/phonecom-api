package main

import (
  "testing"
)

func TestListAccounts(t *testing.T) {

  var app = createCli()

  err := app.Run([]string{"-command", listAccounts})

  if err != nil {
    t.Fatalf("expected no error from Run, got %s", err)
  }

}

