# CoreAPI
## What is it?
This is mostly an example/boilerplate project that can be forked/copied to hook together some of the basic infrastructure of a modern go web API
In the future, I'd like to make this a generator that can create projects in the format laid out. Currently this is very early WIP, and only contains
Basic structure. 

This is also used to test and build out stuff over at [andyantrim/apiutil](https://github.com/andyantrim/apiutil), which is mainly my focus Currently.

## Layout
```
➜  coreapi git:(main) tree . -I vendor
.
├── cmd
│   └── coreapi
│       └── main.go         # Main entry point to the app, more entry points can be added as needed
├── connectors              # Connectors, where each sub package is the service to connect to
│   └── mysql               # WIP: sql connection
│       └── resource.go
├── entity                  # Entity is the core data models of the application
│   ├── base.go
│   └── resource.go
├── go.mod
├── go.sum
├── service                # Service is where the main logic of the app should take place
│   ├── resource.go
│   └── service.go
└── transport             # Transport handles anything coming over the web, this will include gRPC
    ├── http              # http is where the API controllers live, these should only handle request logic.
    │   └── resource.go
    └── router.go
```
