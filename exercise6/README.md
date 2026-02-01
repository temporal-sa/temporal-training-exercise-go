# Exercise 6: Complete Implementation

## Overview
This exercise contains the complete implementation of all features from exercises 1-5:
- Multiple activities with error simulation
- Signal handling for approval
- Query handlers for status tracking
- Search attributes for workflow filtering
- Activity summaries for observability

## Running the Exercise

1. Start the Temporal server:
```bash
temporal server start-dev --search-attribute AccountId=Keyword
```

2. Run the worker:
```bash
go run exercise6/worker/main.go
```

3. Execute the workflow:
```bash
go run exercise6/starter/main.go
```

4. Check the Temporal Web UI at http://localhost:8233

## Features Demonstrated
- Complete money transfer workflow
- Signal-based approval mechanism
- Query-based status monitoring
- Search attributes for filtering
- Activity summaries for observability
