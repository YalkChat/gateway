module gateway

go 1.21.3

require (
	github.com/AleRosmo/cattp v0.0.0-20230903041252-1df1a57e3df1
	github.com/AleRosmo/engine v0.0.0-20231011234010-6038549a07fd
	github.com/AleRosmo/shared_errors v0.0.0-20231011234210-ab222174ab1a
	github.com/YalkChat/database v0.0.0-20231023223418-a7ff42c572b8
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.14.0
	gorm.io/gorm v1.25.5
	nhooyr.io/websocket v1.8.9
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/rs/cors v1.10.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gorm.io/driver/postgres v1.5.3 // indirect
)

replace github.com/YalkChat/database => ../database
