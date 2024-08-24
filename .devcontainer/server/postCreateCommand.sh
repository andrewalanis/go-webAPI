go version
go install github.com/air-verse/air@latest
cp go.mod go.sum ./
go mod download
air -c .air.toml