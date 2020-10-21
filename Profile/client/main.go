package main

import (
	dao "Profile/client/Database"
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	var reply dao.Profile
	var profiles []dao.Profile

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	yoogesh := dao.Profile{Username: "ysharma@gmail.com", FirstName: "Yoogesh", LastName: "Sharma"}
	sushila := dao.Profile{Username: "sushila@gmail.com", FirstName: "Sushila", LastName: "Sapkota"}
	kristy := dao.Profile{Username: "kristy@gmail.com", FirstName: "Kristy", LastName: "Sharma"}

	client.Call("API.CreateProfile", yoogesh, &reply)
	client.Call("API.CreateProfile", sushila, &reply)
	client.Call("API.CreateProfile", kristy, &reply)

	kristy.FirstName = "Krisha"
	client.Call("API.UpdateProfile", kristy, &reply)
	client.Call("API.DeleteProfile", "kristy@gmail.com", &reply)

	client.Call("API.GetAllProfiles", "", &profiles)
	fmt.Println("All profiles: ", profiles)
}
