package main

import (
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
	"github.com/fsouza/go-dockerclient"
	"os"
	"strings"
)

func main() {
	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		panic(err.Error())
	}

	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
	fs := boshsys.NewOsFileSystem(logger)

	// imageOptions := docker.ImportImageOptions{
	// 	Repository: "bosh:stemcell",
	// 	Source:     "/home/vagrant/stemcell/image",
	// }

	// err = client.ImportImage(imageOptions)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// 	settingsDir, err := fs.TempDir("docker-cpi")
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	defer fs.RemoveAll(settingsDir)

	// 	hostConfig := &docker.HostConfig{
	// 		Privileged:      true,
	// 		PublishAllPorts: true,
	// 	}
	// 	containerOptions := docker.CreateContainerOptions{
	// 		Config: &docker.Config{
	// 			ExposedPorts: map[docker.Port]struct{}{
	// 				docker.Port("6868"): {},
	// 			},
	// 			Image: "bosh:stemcell",
	// 			Cmd: []string{
	// 				"/usr/sbin/runsvdir-start",
	// 			},
	// 		},
	// 		HostConfig: hostConfig,
	// 	}

	// 	container, err := client.CreateContainer(containerOptions)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	err = client.StartContainer(container.ID, hostConfig)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	userData := `
	// {
	//   "server": {
	//     "name": "instance-id"
	//   },
	//   "registry": {
	//     "endpoint": "/var/vcap/bosh/settings_source.json"
	//   }
	// }
	// 	`

	// 	settings := `
	// {
	// 	"agent_id": "agent-id",
	// 	"vm": {
	// 		"name": "vm-name",
	// 		"id": "vm-id"
	// 	},
	// 	"mbus": "https://user:password@0.0.0.0:6868/agent",
	// 	"blobstore": {
	//     	"provider": "local",
	//     	"options": {
	//       		"blobstore_path": "/var/vcap/data/blobstore"
	//     	}
	//   	}
	// }
	// 	`

	// 	exec, err := client.CreateExec(docker.CreateExecOptions{
	// 		Container:   container.ID,
	// 		AttachStdin: true,
	// 		Cmd:         []string{"sh", "-c", "cat $1 > /var/vcap/bosh/warden-cpi-user-data.json"},
	// 	})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	err = client.StartExec(exec.ID, docker.StartExecOptions{
	// 		InputStream: strings.NewReader(userData),
	// 	})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	exec, err = client.CreateExec(docker.CreateExecOptions{
	// 		Container:   container.ID,
	// 		AttachStdin: true,
	// 		Cmd:         []string{"sh", "-c", "cat $1 > /var/vcap/bosh/settings_source.json"},
	// 	})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	err = client.StartExec(exec.ID, docker.StartExecOptions{
	// 		InputStream: strings.NewReader(settings),
	// 	})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// println(container.ID)
}
