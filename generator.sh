mvn org.apache.maven.plugins:maven-dependency-plugin:2.8:get -DremoteRepositories=https://mvnrepository.com/ -Dartifact=io.swagger:swagger-codegen-cli:LATEST -Dpackaging=jar -Ddest=swagger-codegen.jar
java -jar swagger-codegen.jar generate -i phonecom.json -l android -o SDKs/android-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l csharp -o SDKs/csharp-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l go -o SDKs/go-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l java -o SDKs/java-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l objc -o SDKs/objc-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l php -o SDKs/php-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l python -o SDKs/python-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l ruby -o SDKs/ruby-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l swift3 -o SDKs/swift3-client-generated
java -jar swagger-codegen.jar generate -i phonecom.json -l typescript-node -o SDKs/typescript-node-client-generated
