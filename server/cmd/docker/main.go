package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"time"
)

func main() {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	res, err := c.ContainerCreate(
		ctx,
		&container.Config{
			Image: "mongo",
			ExposedPorts: nat.PortSet{
				"27017/tcp": {},
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				"27017/tcp": []nat.PortBinding{
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
	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("container started")
	time.Sleep(5 * time.Second)
	fmt.Println("killing container")
	//c.ContainerKill()
	//c.ContainerStop()
	err = c.ContainerRemove(ctx, res.ID, types.ContainerRemoveOptions{
		//Force: true, // force remove
	})
	if err != nil {
		panic(err)
	}
}
