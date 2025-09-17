package main

type User struct {
	Membership
	Name string
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	// code here
	if membershipType == "premium" {
		return User{
			Name: name,
			Membership: Membership{
				Type:             "premium",
				MessageCharLimit: 1000,
			},
		}
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             "standard",
			MessageCharLimit: 100,
		},
	}
}
