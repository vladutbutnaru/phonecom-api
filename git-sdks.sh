#!/usr/bin/env bash

if [[ $# -lt 1 ]] ; then
  echo 'Please enter path to the SDKs folder'
  exit 0
fi

sdkDir=$1

declare -a sdks=(
  "go"
  "android"
  "csharp"
  "java"
  "objc"
  "php"
  "python"
  "ruby"
  "swift3"
  "typescript-node"
)

for sdk in "${sdks[@]}"
do
  dir=$sdkDir/$sdk-client
  echo "Checking status for client: $sdk"
  git --git-dir=$dir/.git/ --work-tree=$dir/ status
done
