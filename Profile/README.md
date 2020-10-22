## Steps to follow

> Go to https://developers.google.com/protocol-buffers/ and click to the download section </br>
> The above steps will bring you to the github page. Go to releases and download the release package based on your operating system </br>
> Unzip the file and copy the "protoc" file from the bin folder and paste it inside our profile folder </br>
> Copy all the file from the "include" folder inside the "third_party" folder which is in our profile folder </br>
> From the terminal, execute below command: </br>
&nbsp;&nbsp;&nbsp;        > go get -u google.golang.org/grpc </br>
&nbsp;&nbsp;&nbsp;        > go get -u github.com/golang/protobuf/protoc-gen-go </br>

> GO inside "go/bin" folder (in the case of window C:\Users\yooge\go\bin) and copy "proto-gen-go" file inside our profile folder here </br>
> Now, run below command to generate a proto file: </br>
&nbsp;&nbsp;&nbsp; protoc --proto_path=proto --proto_path=Mpath/third_party --go_out=plugins=grpc:proto service.proto


