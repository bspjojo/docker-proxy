package networkdata

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func GetNetworksForContainer(container *docker.Container) []string {
	fmt.Println(container.Name + ": " + container.NetworkSettings.NetworkID)

	length := len(container.NetworkSettings.Networks)
	rVals := make([]string, length)
	idx := 0

	for k := range container.NetworkSettings.Networks {
		rVals[idx] = k

		idx = idx + 1
	}

	return rVals
}
