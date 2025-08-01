# gin-basic-structure
The project uses the Gin web framework for building efficient HTTP services. Wire is adopted for automatic dependency injection. GORM handles database interactions, while Redis is used for caching to boost performance and scalability.

## Project Structure
```
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── config                    # All configuration files
│   └── config.example.yaml
├── internal
│   ├── conf                  # Configuration loaders and structs
│   │   ├── conf.go           # Load config files
│   │   ├── data.go           # Structs for DB and Redis configs
│   │   └── server.go         # Struct for server settings
│   ├── controller            # Controllers
│   │   ├── index.go
│   │   └── provider.go       # Controller providers
│   ├── dao                   # Data access layer
│   │   ├── provider.go       # dao providers
│   │   └── user.go
│   ├── data                  # DB and Redis connections
│   │   └── data.go
│   ├── entity                # Database entity definitions
│   │   └── users.gen.go
│   ├── handler               # HTTP response handling
│   │   └── response.go
│   ├── middleware            # Middlewares
│   │   └── auth.go
│   ├── router
│   │   ├── api_router.go
│   │   └── provider.go       # router providers
│   ├── schema                # Request/response schemas
│   │   └── user.go
│   ├── server                # HTTP server setup
│   │   └── http.go
│   └── service               # Business logic services
│       ├── provider.go       # service providers
│       └── user.go
└── utils                     # Some local utils functions saved in here
```


## License

[MIT license](https://github.com/zhaoxlchn/gin-basic-structure/blob/main/LICENSE)