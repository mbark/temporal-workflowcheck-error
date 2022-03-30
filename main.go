package temporal_template_workflowcheck

import (
	temporalClient "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	clientOptions := temporalClient.Options{}
	cli, err := temporalClient.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("failed to setup client: %s", err)
	}
	defer cli.Close()

	w := worker.New(cli, "TaskQueue", worker.Options{
		EnableLoggingInReplay: true,
	})
	w.RegisterWorkflow(Workflow)
}
