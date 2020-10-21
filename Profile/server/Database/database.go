package Database

type Profile struct {
	Username  string
	FirstName string
	LastName  string
}

type API int

var database []Profile

func (a *API) CreateProfile(newProfile Profile, reply *Profile) error {
	database = append(database, newProfile)
	*reply = newProfile
	return nil
}

func (a *API) UpdateProfile(newProfile Profile, reply *Profile) error {
	var updatedProfile Profile
	for k, v := range database {
		if v.Username == newProfile.Username {
			database[k] = Profile{Username: newProfile.Username, FirstName: newProfile.FirstName, LastName: newProfile.LastName}
			updatedProfile = database[k]
		}
	}
	*reply = updatedProfile
	return nil
}

func (a *API) DeleteProfile(username string, reply *Profile) error {
	var detetedProfile Profile

	for k, v := range database {
		if v.Username == username {
			detetedProfile = Profile{Username: v.Username, FirstName: v.FirstName, LastName: v.LastName}
			database = append(database[:k], database[k+1:]...)
			break
		}
	}
	*reply = detetedProfile
	return nil
}

func (a *API) GetProfileByUsername(Username string, reply *Profile) error {
	var profile Profile

	for _, v := range database {
		if v.Username == Username {
			profile = v
		}
	}
	*reply = profile
	return nil
}

func (a *API) GetAllProfiles(empty string, reply *[]Profile) error {
	*reply = database
	return nil
}
