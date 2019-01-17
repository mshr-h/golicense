# golicense

Command line tool to generate LICENSE file

## Installation

```
go get github.com/mshr-h/golicense
```

## Usage

```
$ golicense
Invalid license type
  -author string
    	Author name (default "none")
  -license string
    	License type (default "unknown")
```

Generate `MIT` LICENSE which author is `mshr-h`.

```
$ golicense -author mshr-h -license MIT
```

## Supported license types

Currently these license types are supported.

- MIT
- Apache 2.0
- GPL-3
