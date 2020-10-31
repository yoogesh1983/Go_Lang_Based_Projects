## Steps to follow

- Go to https://developers.google.com/protocol-buffers/ and click to the download section
- The above steps will bring you to the github page. Go to releases and download the release package based on your operating system
- Unzip the file and copy the "protoc" file from the bin folder and paste it inside our profile folder
- Copy all the file from the "include" folder inside the "third_party" folder which is in our profile folder
- From the terminal, execute below command: </br>

    > go get -u google.golang.org/grpc </br>
    > go get -u github.com/golang/protobuf/protoc-gen-go </br>

- Go inside the "go/bin" folder (in the case of window C:\Users\yooge\go\bin) and copy "proto-gen-go" file inside our profile folder here
- Now, run below command to generate a proto file: </br>
    > protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto


