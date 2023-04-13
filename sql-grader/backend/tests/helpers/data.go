package helpers

import "backend/utils/value"

var D = &TestData{
	Users: []*User{
		{
			Uid:  value.Ptr("ajZYATlqKPGgCgO9QnwURx20vDuw"),
			Name: value.Ptr("Peach Mountain"),
		},
		{
			Uid:  value.Ptr("k8lWGRyeSijkXpmM2rQImwXg5zLI"),
			Name: value.Ptr("Raccoon Olive"),
		},
		{
			Uid:  value.Ptr("uFJWrDJV7npfQRTII4RvlXLD78JL"),
			Name: value.Ptr("Panda Algae"),
		},
	},
}

type TestData struct {
	Users []*User `json:"users"`
}

type User struct {
	Uid   *string `json:"uid"`
	Token *string `json:"token"`
	Name  *string `json:"name"`
}
