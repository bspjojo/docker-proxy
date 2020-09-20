package handler

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	configTypes "main/configStructs"
	"main/handler/stringBuilder"
	"main/handler/writer"

	docker "github.com/fsouza/go-dockerclient"
)

func HandleDockerEvent(msg *docker.APIEvents, client *docker.Client, config configTypes.AppConfig) {
	if msg.Type == "container" {
		log.Println("=============================")
		log.Println(msg)
		image := msg.Actor.Attributes["image"]
		log.Println("image ", image)

		switch msg.Action {
		case "create",
			"destroy",
			"die",
			"exec_create",
			"exec_detach",
			"exec_die",
			"exec_start",
			"kill",
			"pause",
			"restart",
			"start",
			"stop",
			"unpause",
			"update":
			log.Println(msg.Action)

			log.Println("Action ", msg.Action)
			log.Println("Type ", msg.Type)
			log.Println("Actor ", msg.Actor)
			log.Println("Status ", msg.Status)
			log.Println("From ", msg.From)
			log.Println("Time", msg.Time)
			log.Println("TimeNano", msg.TimeNano)
			log.Println("=============================")

			OrchestrateBuildConfig(client, config)
		}
	}
}

func OrchestrateBuildConfig(client *docker.Client, config configTypes.AppConfig) {
	content := stringBuilder.BuildString(client)
	writer.WriteContent(content, config.OutPath)

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		reloadConfig()
	}
}

func reloadConfig() {
	out, err := exec.Command("nginx", "-s", "reload").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
