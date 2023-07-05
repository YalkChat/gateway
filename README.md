# What is this?
It's a simple chat application backend, it's meant to provide a connection gateway to interact with in order to communicate with other clients on the network.
As of now, it's meant to be self-deployed on your premises or in the cloud 

# This looks like Discord!
Yeah of course I'm literally copying and recreating from scratch what they do, a succesful product it's always an example to follow and their documentation is simply phenomenal.
No, I don't want to copy Discord, I am just learning from their approach


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

# Disclaimer
This is by no means close to completion and has an incredible amount of security flaws and bad programming practices, **do not use this in production!** (it doesn't even work for most parts, requires a lot of manual steps which will be all addressed overtime)

Massive thanks to the entirety of the Open Source communities of the world. Freedom of Information is what fueled my learning journey and I will never be enough thankful to all of you who decided one day to write an article, make a Medium post, ask a question on StackOverflow, etc. that I have stumbled upon and that I've learnt so much or solved a problem. Seriously, thank you.

As for the rest take this code and do whatever you want, maybe mention me if you use it and feel free to do whatever you want with it but it's always a work of the Open Source communities. Also, feel free to open any PR or comment, I am glad to hear anyone's opinion.

# Why am I doing this?
It's my learning journey into DevOps and in building a complete product from ground up, sticking to the least amount of dependencies possible to face those challenges that enables you to properly learn what is the right thing to do and to dive as deep as possible in GoLang.
This was something I always dreamed to do and only recently gained the skills and experience to actually attempt to create this, even more important I've learnt how to learn and to overcome challenges I could've never attempted to solve before.
Building this entire application by myself has taught and keep teaching me new things crossing the boundaries of the entire Tech spectrum. Most importantly: I'm loving every part of this journey and having incredible fun.
