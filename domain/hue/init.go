package hue

import (
	"github.com/heatxsink/go-hue/groups"
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
func findCorrectGroup(name string) (*groups.Group, error) {
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
func getAllGroup() ([]groups.Group, error) {
	groups, err := getGroupClient().GetAllGroups()
	if err != nil {
		return nil, err
	}

	return groups, nil
}
