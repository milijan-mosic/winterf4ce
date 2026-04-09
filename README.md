# Winterf4ce

## Dev setup

### Clone project

```sh
git clone https://github.com/windwalk-bushido/winterf4ce
cd winterf4ce/
```

### Run it

```sh
air server --port 20000 -c .air.toml
```

## Prod setup

### Build

```sh
docker build -t winterf4ce .
```

### Run it

```sh
docker run -p 20000:20000 winterf4ce
```

## How to write/edit websites?

Open up any file in "data" directory and write it in format:

- name
- URL
- _empty row_
