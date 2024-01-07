# FreeBSD Go Lang Reverse PRoxy

HTTP Web Server

## Features

- Reverse Proxying
    - Load Balancing
    - Active Healthchecks
- File Server
    - Browsing Directories
    - Serving Static Sites
- HTTP Response Header Modification
- Authentication
    - Basic HTTP Auth
    - Key Auth
- Optional TLS Support
- URL Rewrites

## Getting Started

**Pre-requisites**
- [Go >= 1.18](https://go.dev/)
- [GNU Make](https://www.gnu.org/software/make/)

### Configuration
- See [CONFIG.md](./CONFIG.md)

### Build from Source

git clone https://github.com/unixwinbsd/Go-ReverseProxy-FreeBSD.git
go mod init github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD
go mod tidy

go get golang.org/x/crypto/bcrypt
go get github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD/internal/reverseproxy
go get github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD
go get github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD/internal/server
go get github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD/internal/config

The final step is to run the "make build" command, to create a reverse proxy bin file.
make build

This will create a binary in the `bin` directory.

### Usage

```shell
./unixwinbsd
```
Default config file path: `./config.json`



## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.

## Acknowledgements

- Full Article
  - https://www.unixwinbsd.site
