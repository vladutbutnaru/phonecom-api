#!/bin/bash

swaggerFilePath=phonecom.json

mvn org.apache.maven.plugins:maven-dependency-plugin:2.10:get -DremoteRepositories=central -Dartifact=io.swagger:swagger-codegen-cli:LATEST -Dpackaging=jar -Ddest=swagger-codegen.jar

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

sdkDir=SDKs

rm -r $sdkDir

for sdk in "${sdks[@]}"
do
  java -jar swagger-codegen.jar generate -i $swaggerFilePath -l $sdk -o $sdkDir/$sdk-client
done
