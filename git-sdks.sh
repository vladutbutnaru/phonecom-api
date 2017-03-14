#!/usr/bin/env bash

if [[ $# -lt 2 ]] ; then
  echo 'Usage: Please enter path to the SDKs folder and a git command: status or add'
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

if [ $2 == "status" ]
then

  for sdk in "${sdks[@]}"
  do
    dir=$sdkDir/$sdk-client
    echo "Checking status for client: $sdk"
    git --git-dir=$dir/.git/ --work-tree=$dir/ status
  done

fi

if [ $2 == "add" ]
then

  for sdk in "${sdks[@]}"
  do
    dir=$sdkDir/$sdk-client
    echo "Adding all for client: $sdk"
    git --git-dir=$dir/.git/ --work-tree=$dir/ add --all
    git --git-dir=$dir/.git/ --work-tree=$dir/ commit -m "Updating $sdk SDK"
  done

fi

if [ $2 == "push" ]
then

  for sdk in "${sdks[@]}"
  do
    dir=$sdkDir/$sdk-client
    echo "Pushing for client: $sdk"
    echo "git --git-dir=$dir/.git/ --work-tree=$dir/ push origin master"
  done

fi
