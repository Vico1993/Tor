package hue

import (
	"os"

	"github.com/heatxsink/go-hue/groups"
)

func getGroupClient() *groups.Groups {
	return groups.New(
		os.Getenv("HUE_TEST_HOSTNAME"),
		os.Getenv("HUE_TEST_USERNAME"),
	)
}