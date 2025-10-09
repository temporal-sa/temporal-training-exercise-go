# Exercise 1 Solutions

## GreetingWorkflowImpl.java
```java
private final GreetingActivity activity = Workflow.newActivityStub(
    GreetingActivity.class,
    ActivityOptions.newBuilder()
        .setStartToCloseTimeout(Duration.ofSeconds(10))
        .build()
);

@Override
public String greet(String name) {
    return activity.createGreeting(name);
}
```

## GreetingActivityImpl.java
```java
@Override
public String createGreeting(String name) {
    return "Hello, " + name + "!";
}
```

## StartWorker.java
```java
worker.registerWorkflowImplementationTypes(GreetingWorkflowImpl.class);
worker.registerActivitiesImplementations(new GreetingActivityImpl());
```

## Running the Exercise
1. Start worker: `gradle execute -PmainClass=com.temporal.training.exercise1.StartWorker`
2. Execute workflow: `gradle execute -PmainClass=com.temporal.training.exercise1.StartWorkflow`
