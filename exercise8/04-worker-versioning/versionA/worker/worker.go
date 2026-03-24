package main

import (
	"log"
	versiona "versioning/versionA"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "driver-tracking"

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	deploymentOptions := worker.DeploymentOptions{
		UseVersioning: true,
		Version: worker.WorkerDeploymentVersion{
			DeploymentName: "example",
			BuildID:        "version-a",
		},
		DefaultVersioningBehavior: workflow.VersioningBehaviorPinned,
		// DefaultVersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	}

	w := worker.New(c, TaskQueue, worker.Options{
		DeploymentOptions: deploymentOptions,
	})

	// Only one of the following two Workflow registrations can be uncommented
	w.RegisterWorkflow(versiona.TrackDriverLocation)
	// w.RegisterWorkflowWithOptions(versiona.TrackDriverLocation, workflow.RegisterOptions{
	// 	VersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	// })
	w.RegisterActivity(versiona.RecordDriverLocation)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
