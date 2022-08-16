package hue

import (
	"fmt"
	"os"

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

func SetLight(on bool) {
	gg := groups.New(os.Getenv("HUE_TEST_HOSTNAME"), os.Getenv("HUE_TEST_USERNAME"))
	grps, err := gg.GetAllGroups()
	if err != nil {
		panic(err)
	}

	var bri uint8 = 100
	if !on {
		bri = 0
	}

	for _, grp := range grps {
		fmt.Println(grp.ID, grp.Name)

		if (grp.ID == 5) {
			_, err := gg.SetGroupState(grp.ID, lights.State{
				On: on,
				Bri: bri,
				TransitionTime: 1,
			})

			if err != nil {
				panic(err)
			}
		}
	}
}

// func ShutDownLight() {
// 	fmt.Println("Shutting down...")

// 	gg := groups.New(os.Getenv("HUE_TEST_HOSTNAME"), os.Getenv("HUE_TEST_USERNAME"))
// 	grps, err := gg.GetAllGroups()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, grp := range grps {
// 		fmt.Println(grp.ID, grp.Name)

// 		if (grp.ID == 5) {
// 			gg.SetGroupState(grp.ID, lights.State{
// 				On: false,
// 				TransitionTime: 1,
// 			})
// 		}
// 	}
// }