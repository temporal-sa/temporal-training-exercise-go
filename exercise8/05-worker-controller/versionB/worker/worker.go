package main

import (
	"log"
	"os"
	versionb "versioning/versionB"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "driver-tracking"

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "host.minikube.internal:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Worker must use the build ID/deployment name from environment
	// These are set on the deployment by the controller
	buildID := os.Getenv("TEMPORAL_WORKER_BUILD_ID")
	deploymentName := os.Getenv("TEMPORAL_DEPLOYMENT_NAME")
	if buildID == "" || deploymentName == "" {
		log.Fatalln("buildID and/or deploymentName missing")
	}

	deploymentOptions := worker.DeploymentOptions{
		UseVersioning: true,
		Version: worker.WorkerDeploymentVersion{
			DeploymentName: deploymentName,
			BuildID:        buildID,
		},
		// DefaultVersioningBehavior: workflow.VersioningBehaviorPinned,
		DefaultVersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	}

	w := worker.New(c, TaskQueue, worker.Options{
		DeploymentOptions: deploymentOptions,
	})
	w.RegisterWorkflow(versionb.TrackDriverLocation)
	w.RegisterActivity(versionb.RecordDriverLocation)
	w.RegisterActivity(versionb.CheckGeofence)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
