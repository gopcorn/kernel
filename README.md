# gopcorn kernel

### Directory structure:

```sh
.
├── adapters
├── client
├── cmd
│   └── kernel
├── conformance
├── go.mod
├── kernel
├── proto
├── README.md
├── sdk
└── testfeatures

```

| dir          | description                                                              |
| ------------ | ------------------------------------------------------------------------ |
| sdk          | The contract: types and interfaces that features import (own go.mod)     |
| kernel       | Dispatch, host, flags, admin API (own go.mod, imports sdk only)          |
| cmd/kernel   | main.go of the kernel binary. Listens on the Unix socket                 |
| client       | sdk.Host that talks to the kernel over the socket (import sdk and proto) |
| proto        | the socket protocol as protobuf files                                    |
| adapters     | postgres, redis, etc implementation of the sdk ports (own go.mod)        |
| testfeatures | small features used only by tests                                        |
| conformance  | exported test suite that any kernel build has to pass                    |
