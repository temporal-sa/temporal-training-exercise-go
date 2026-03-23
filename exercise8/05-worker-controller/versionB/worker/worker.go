package main

import (
	"log"
	"os"
	exercise8 "versioning/versionB"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

const TaskQueue = "hello-task-queue"

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "host.minikube.internal:7233",
		// HostPort: "localhost:7233",
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
			// DeploymentName: "hello-world/hello-world",
			// BuildID:        "version-b",
			DeploymentName: deploymentName,
			BuildID:        buildID,
		},
		// DefaultVersioningBehavior: workflow.VersioningBehaviorPinned,
		DefaultVersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	}

	w := worker.New(c, TaskQueue, worker.Options{
		DeploymentOptions: deploymentOptions,
	})
	w.RegisterWorkflow(exercise8.Greet)
	w.RegisterActivity(exercise8.CreateGreeting)
	w.RegisterActivity(exercise8.Uppercase)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
