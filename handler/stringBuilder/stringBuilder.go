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
			content, use := buildContainerPart(container)
			if use {
				builder.WriteString("\n")

				builder.WriteString(content)
			}
		}
	}

	builder.WriteString("}")

	return builder.String()
}

func buildContainerPart(container *docker.Container) (string, bool) {
	location, hasLocation := getPathFromContainer(container)

	if !hasLocation {
		return "", false
	}

	var builder strings.Builder
	builder.WriteString("  location " + location + " {\n")

	ports := container.NetworkSettings.Ports

	keys := make([]string, 0, len(ports))
	for k := range ports {
		keys = append(keys, k.Port())
	}

	port := keys[0]

	builder.WriteString("    proxy_pass      http:/" + container.Name + ":" + port + ";\n")
	builder.WriteString("  }\n")

	return builder.String(), true
}

func getPathFromContainer(container *docker.Container) (string, bool) {
	for _, env := range container.Config.Env {
		split := strings.Split(env, "=")
		name := split[0]
		val := split[1]

		if name == "PROXY_LOCATION" {
			fmt.Println(env, name, val)
			return val, true
		}
	}

	return "", false
}
