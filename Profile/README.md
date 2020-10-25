## Environment settins:

- Go to https://developers.google.com/protocol-buffers/ and click to the download section
- The above steps will bring you to the github page. Go to releases and download the release package based on your operating system
- Unzip the file and copy the "protoc" file from the bin folder and put it into Environment path (normally for windows rather than putting it inside
  C:\Users\yooge\go\bin folder,just put inside "C:\go\bin" folder since this path will by default be in an environment path and hence you don't need
  to create a seperate environment path. In mac though, you better put it inside usr/bin)
- Copy all the file from the "include" folder inside the "third_party" folder here in our Profile module
- From the terminal, execute below command: </br>

    > go get -u google.golang.org/grpc </br>
    > go get -u github.com/golang/protobuf/protoc-gen-go </br>

- Go inside the "go/bin" folder in the case of Mac (whereas in the case of window C:\Users\yooge\go\bin) and copy "proto-gen-go"and paste it to an enviromnment path( See above about environment path)

- Now, run below command to generate a proto file: </br>
    > protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto


https://www.youtube.com/watch?v=hjTI35iKMyQ [By directional video]


