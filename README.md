# dts-developer-challenge

## Run
Run both the frontend and backend apps together
Enter the backend/ directory and run the following command from the terminal:  `$ make run`. This will execute the Makefile and will run the backend server.
Enter the frontend directory and run the following command from the terminal:  `$ make run`. This will execute the Makefile and will run the frontend client application.

## Data Format
The data is currently stored in a Mongo DB Atlas online database. The collection contains data in the following format:

```
{
  "id": "string",
  "title": "string", 
  "description": "string",
  "status": "pending|in-progress|completed",
  "due_date": "2025-06-24T10:00:00Z"
}
```

## API Documentation 
### Swagger Docs
When running the backend API, go to: http://localhost:8080/swagger/index.html

## What I learned
- No  trailing "/" in routes otherwise this will lead to CORS errors when attempting to connect to the endpoint from the frontend application