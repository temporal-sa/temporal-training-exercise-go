# Exercise 1 Solutions

## workflow.go
```go
func (w *GreetingWorkflow) Greet(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * workflow.Second}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, CreateGreeting, name).Get(ctx, &result)
	return result, err
}
```

## activity.go
```go
func CreateGreeting(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
```

## worker/main.go
```go
w.RegisterWorkflow((&exercise1.GreetingWorkflow{}).Greet)
w.RegisterActivity(exercise1.CreateGreeting)
```

## Running the Exercise
1. Start worker: `go run exercise1/worker/main.go`
2. Execute workflow: `go run exercise1/starter/main.go`
