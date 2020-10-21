package Database

import "fmt"

type Profile struct {
	Username  string
	FirstName string
	LastName  string
}

var database []Profile

//@init
func Init() {
	CreateProfile(Profile{Username: "ysharma@gmail.com", FirstName: "Yoogesh", LastName: "Sharma"})
	CreateProfile(Profile{Username: "sushila@gmail.com", FirstName: "Sushila", LastName: "Sapkota"})
	CreateProfile(Profile{Username: "kristy@gmail.com", FirstName: "kristy", LastName: "Sharma"})
	fmt.Println("Database initialized: ", *GetAllProfiles())
}

func CreateProfile(newProfile Profile) ([]Profile, error) {
	database = append(database, newProfile)
	return database, nil
}

func UpdateProfile(newProfile Profile) (Profile, error) {
	var updatedProfile Profile
	for k, v := range database {
		if v.Username == newProfile.Username {
			database[k] = Profile{Username: newProfile.Username, FirstName: newProfile.FirstName, LastName: newProfile.LastName}
			updatedProfile = database[k]
		}
	}
	return updatedProfile, nil
}

func DeleteProfile(username string) (Profile, error) {
	var detetedProfile Profile

	for k, v := range database {
		if v.Username == username {
			detetedProfile = Profile{Username: v.Username, FirstName: v.FirstName, LastName: v.LastName}
			database = append(database[:k], database[k+1:]...)
			break
		}
	}
	return detetedProfile, nil
}

func GetProfileByUsername(Username string) (Profile, error) {
	var profile Profile

	for _, v := range database {
		if v.Username == Username {
			profile = v
		}
	}
	return profile, nil
}

func GetAllProfiles() *[]Profile {
	return &database
}
