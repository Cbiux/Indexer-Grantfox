# Trustless Work Indexer

Blockchain indexer for the Stellar/Soroban network. Processes ledgers in real-time, extracting transactions, operations, and state changes.

## Prerequisites

- **Go 1.25+** - [Download](https://go.dev/dl/)
- **Make** (optional) - Usually pre-installed on macOS/Linux

### Verify installation

```bash
go version
# Should display: go version go1.25.x or higher
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Trustless-Work/Indexer.git
cd Indexer
```

2. Download dependencies:

```bash
go mod download
```

## Running the project

### Option 1: Using Make

```bash
# Build and run
make run
```

### Option 2: Go commands

```bash
# Build
go build -o bin/ingest cmd/ingest.go

# Run
./bin/ingest
```

### Option 3: Direct execution (without building)

```bash
go run cmd/ingest.go
```

## Configuration

The indexer is configured by default to connect to **Stellar Testnet**:

| Parameter | Value |
|-----------|-------|
| Network | Stellar Testnet |
| RPC URL | `https://soroban-testnet.stellar.org` |

To modify the configuration, edit the constants in `cmd/ingest.go`.

## Project structure

```
.
├── cmd/
│   └── ingest.go          # Entry point
├── internal/
│   ├── indexer/           # Processing engine
│   ├── ingest/            # Ingestion configuration
│   ├── services/          # RPC services
│   └── entities/          # Data structures
├── bin/                   # Compiled binaries
├── Makefile
└── go.mod
```

## Useful commands

```bash
# Build
make build

# Build and run
make run

# Clean binaries
rm -rf bin/
```