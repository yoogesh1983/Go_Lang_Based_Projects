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

## How to connect from client side

 - go inside the client folder and give below command. If you don'y proivde <strong>-N {name}</strong>, then it will by default use the name as <strong>kristy</strong> that we defined at transport.go file. it will then generate the id as the hash of the given name:

    > go run main.go -N yoogesh

- You can create as many as client you want.
 - Now you can enter whatever the message you want from the console and when done click enter. upon doing that, it will return you a hashed value of id along with the message you sent. Now if you go to other client console, you can see the same message is broadcast there:
 
    > yooge@DESKTOP-B96F47M MINGW64 ~/Documents/GitHub/Go_Lang_Based_Projects/Profile/client (08-Grpc_part2_chatApplication)
    > $ go run main.go -N yoogesh
    > Good morning!!!
    > 1b05a7327525fd65e87df3494f608bf3f7dea2ac2c6d7266d6c4f6115069638f : Good morning!!





