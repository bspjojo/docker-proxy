package stringBuilder

import (
	"fmt"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

func BuildString(client *docker.Client) string {
	containers, err := client.ListContainers(docker.ListContainersOptions{true, false, 0, "", "", nil, nil})

	if err != nil {
		fmt.Println(err)
		return "err"
	}

	var builder strings.Builder

	builder.WriteString("server {\n")
	builder.WriteString("  listen       80;\n")

	for _, apiContainer := range containers {
		container, err := client.InspectContainer(apiContainer.ID)

		if err != nil {
			fmt.Println(err)
			return "err"
		}

		if container.State.Running {
			builder.WriteString("\n")

			builder.WriteString(buildContainerPart(container))
		}
	}

	builder.WriteString("}")

	return builder.String()
}

func buildContainerPart(container *docker.Container) string {
	var builder strings.Builder
	builder.WriteString("  location /proxy" + container.Name + " {\n")

	ports := container.NetworkSettings.Ports

	keys := make([]string, 0, len(ports))
	for k := range ports {
		keys = append(keys, k.Port())
	}

	port := keys[0]

	builder.WriteString("    proxy_pass      http:/" + container.Name + ":" + port + ";\n")
	builder.WriteString("  }\n")

	return builder.String()
}
