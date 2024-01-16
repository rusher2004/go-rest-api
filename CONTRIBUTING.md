# Contributing

## Adding an Endpoint

To add an endpoint to the server, make the following updates:

1. Update the [OpenAPI spec file](./openapi.yaml). Every operation must include expected input parameters and response schemas.
2. Create the handler function to be called by the HTTP router
   - This is a standard [http.handler](https://pkg.go.dev/net/http#Handler.ServeHTTP) func, implemented as a method on the `Server` struct. It will validate user input and call the relevant data service method to handle the business logic of the request.
3. Server Route
   - Update the [server routes](./server/server.go?plain=1#L70) in `(s *Server) Routes()` to include the path to the new endpoint, with the new handler func assigned.
   - Be sure to properly apply any necessary middleware here.
4. Update DataStore definition
   - Add the function to be utilized by your handler to the [DataStore](ecs/api-router/server/server.go?plain=1#L12) interface.
   - Each function should have an input and output struct defined, the exception being `DELETE` endpoints, where no output is expected.
5. Add the new function definition to the [New DataStore](./new-store/) and [Old DataStore](./old-store/store.go) directory.
   - This is where you will implement the business logic to statisfy the request according to what the processor requires.
   - If one of the processors does not support the method, return a `server.ErrNotImplemented` error.
6. Add integration/unit tests
   - Validate input (processor, parameter format validation, etc.).
   - Validate happy path for each endpoint.
   - Test that responses from the different data sources result in the expected errors/responses.
