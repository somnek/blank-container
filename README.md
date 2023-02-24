# blank-container
Quickest way to spin up empty Docker container(s) 🧩

## Installation
Only support git clone for now.

### Git Clone
```
git clone https://github.com/somnek/blank-container.git
```

## Usage
```bash
    ./whale [OPTIONS]
        --up       Run empty container
        --clean    Remove the image & container
```

## Example
* create multiple containers
```bash
    ./whale up --count=4
```

This example will create 4 containers
