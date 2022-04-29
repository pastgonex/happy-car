package mongotesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"testing"
)

const (
	image         = "mongo"
	containerPort = "27017/tcp"
)

// RunWithMongoInDocker runs the tests with
// a mongodb instance in a docker container.
func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	// start container
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	res, err := c.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			ExposedPorts: nat.PortSet{
				containerPort: {},
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				containerPort: []nat.PortBinding{
					{
						HostIP: "127.0.0.1",
						//HostPort: "27018",
						HostPort: "0", // auto-allocate port
					},
				},
			},
		}, nil, nil, "test_mongo")
	if err != nil {
		panic(err)
	}
	containerID := res.ID
	defer func() {
		//c.ContainerKill()
		//c.ContainerStop()
		err = c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true, // force remove
		})
		if err != nil {
			panic(err)
		}
	}()

	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	inspectResponse, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	hostPort := inspectResponse.NetworkSettings.Ports[containerPort][0]
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	return m.Run()
}
