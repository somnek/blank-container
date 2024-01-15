# blank-container

Quickest way to spin up empty Docker container(s) 🧩

## Installation

Only support git clone for now.

### Clone & build

```
git clone https://github.com/somnek/blank-container.git
go build -o blank .
```

then you can move the binary to your $PATH

## Usage

```bash
    ./blank [OPTIONS]
        --up       Run empty container
        --clean    Remove the image & container
```

## Example

- create multiple containers

```bash
    ./blank up --count=4
```

This example will create 4 containers
