# Exercise 4 Solution: Visibility & Monitoring with Custom Search Attributes

## Key Implementation Points

### 1. Import SearchAttributeKey
```java
import io.temporal.common.SearchAttributeKey;
```
- Required for type-safe search attribute operations

### 2. Upsert Search Attributes in Workflow
```java
@Override
public String transfer(TransferRequest request) {
    this.request = request;
    logger.info("Starting transfer: {}", request);
    
    // Set search attributes for filtering
    Workflow.upsertTypedSearchAttributes(
        SearchAttributeKey.forKeyword("AccountId").valueSet(request.fromAccount())
    );
    
    // ... rest of workflow logic
}
```

### 3. Key Components Explained

**SearchAttributeKey.forKeyword("AccountId")**
- Creates a keyword-type search attribute key
- "AccountId" must match the attribute defined when starting Temporal server
- Keyword attributes are for exact text matching

**valueSet(request.fromAccount())**
- Sets the search attribute value to the source account ID
- Enables filtering workflows by account

**Workflow.upsertTypedSearchAttributes()**
- Updates search attributes for the current workflow execution
- Can be called multiple times to update values
- Changes are immediately visible in Temporal Web UI

## Key Concepts Demonstrated

1. **Custom Search Attributes**: Enable workflow filtering and discovery
2. **Type Safety**: SearchAttributeKey provides compile-time type checking
3. **Runtime Updates**: Search attributes can be set/updated during workflow execution
4. **Workflow Visibility**: Enhanced monitoring and debugging capabilities

## Testing the Solution

1. **Start Worker**: `./gradlew run -PmainClass=com.temporal.training.solution4.StartWorker`
2. **Run Workflow**: `./gradlew run -PmainClass=com.temporal.training.solution4.StartWorkflow`
3. **Verify in Web UI**: Open http://localhost:8233 and check workflow details

## Expected Results
- Workflow executes with all previous functionality intact
- AccountId search attribute visible in Temporal Web UI
- Can filter workflows using: `AccountId = "account-123"`
- Search attribute appears in workflow execution details

## Search Attribute Types
- **Keyword**: Exact text matching (used in this exercise)
- **Text**: Full-text search capabilities
- **Int**: Numeric values with range queries
- **Double**: Floating-point numbers
- **Bool**: Boolean values
- **Datetime**: Date and time values