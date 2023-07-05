# What is this?
It's a simple chat application backend

# How do I run this?
* Get a fresh Postgres instance
* Clone the repo
* Configure the .env file
* Open a terminal in the folder 
* Execute with ``go run .``

## To-Do
- [ ] Accept tokens in RawEvent payload (Identify)
- [ ] Integrating heartbeat for status detection
- [ ] Include user automatically in General chat
- [ ] Sequential messages in a 5m timestamp shows only as timestamp (no pfp, no name)

## Known bugs
- [ ] After login redirect doesn't work
- [ ] Initial DB configuration breaks the app
- [ ] Root path doesn't redirect anywhere
- [ ] If user is in no chat, it just breaks
- [ ] Toast component shows for everyone, should be only admins
- [ ] New Channel/DM button opens new user modal

## File Structure in transition
I'm moving to a more meaningful file structure, the change is driven by the need of separating the concerns and break down the monolithic chat server core into a more modular solution.
This is the proposal 

```
yalk-backend/
├── cmd/                  # Application entry points
│   └── yalk/             # Main application
│       └── main.go       # Main program
├── pkg/                  # Library code
│   ├── chat/             # Chat-related functionality
│   │   ├── service.go    # Business logic for chats
│   │   └── repository.go # Database access for chats
│   ├── user/             # User-related functionality
│   │   ├── service.go    # Business logic for users
│   │   └── repository.go # Database access for users
│   ├── auth/             # Authentication functionality
│   │   ├── service.go    # Business logic for authentication
│   │   └── repository.go # Database access for authentication
│   └── ...               # Other packages
├── api/                  # API handlers
│   ├── chat/             # Handlers for chat-related endpoints
│   ├── user/             # Handlers for user-related endpoints
│   ├── auth/             # Handlers for authentication endpoints
│   └── ...               # Other handlers
├── test/                 # Test files
├── scripts/              # Scripts for tasks like building or deploying
├── Dockerfile            # Docker configuration
├── .env                  # Environment variables
└── ...                   # Other files (README, .gitignore, etc.)
```