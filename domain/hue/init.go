package hue

import (
	"fmt"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
)

/*
	l := lightAPI.New(os.Getenv("HUE_TEST_HOSTNAME"), os.Getenv("HUE_TEST_USERNAME"))
	lights, err := l.GetAllLights()
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of lights", len(lights))

	for _, light := range lights {
		fmt.Println(light.ID,light.Name)

		if light.ID == 14 {
			res, err := l.SetLightState(light.ID, lightAPI.State{
				On: false,
				TransitionTime: 1000,
			})
			if err != nil {
				panic(err)
			}

			fmt.Println(res)
		}
	}
*/

// Get a group based on a Name
func FindCorrectGroup(name string) (*groups.Group, error) {
	groups, err := getGroupClient().GetAllGroups()
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		if group.Name == name {
			return &group, nil
		}
	}

	return nil, nil
}

// Get all Groups for your hue
func GetAllGroup() ([]groups.Group, error) {
	groups, err := getGroupClient().GetAllGroups()
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func ShutDownGroup(name string) {
	group, err := FindCorrectGroup(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("SHUTDOWN")

	_, err = getGroupClient().SetGroupState(
		group.ID,
		lights.State{
			On: false,
			Bri: 0,
			TransitionTime: 1,
		},
	)

	if err != nil {
		panic(err)
	}
}

func PowerUpGroup(name string) {
	group, err := FindCorrectGroup(name)
	if err != nil {
		panic(err)
	}

	_, err = getGroupClient().SetGroupState(
		group.ID,
		lights.State{
			On: true,
			Bri: 200,
			TransitionTime: 1,
		},
	)

	if err != nil {
		panic(err)
	}
}
