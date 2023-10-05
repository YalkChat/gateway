The handlers use the abstractions on the server, which operates with the database within it's method and never exposing the DB models and implementations.
We don't do the same with the sessionManager because the chat server doesn't care about HTTP sessions.
We could consider instead to just add the config in th server type with an interface.

There might be a lot of repetitions in the handlers and it needs to be tackled


## Future Refactoring Considerations

### Authentication Logic
- The authentication logic is currently implemented directly within the `signin` handler.
- Consider moving this logic into its own subpackage for better separation of concerns and reusability.
- This will become particularly important as the project grows and potentially needs to handle authentication in multiple places.
