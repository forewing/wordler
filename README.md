# WORDLER

## Command Line

### Usage

```
./wordler -h
Usage of ./wordler:
  -at {index}:{letter}
        comma-seperated {index}:{letter}, where {letter} is at {index} (0-based) of the target
  -has string
        comma-seperated strings the target contains
  -len int
        length of word
  -max int
        max output words (default 20)
  -no string
        comma-seperated strings the target does not contain
  -not-at {index}:{letter}
        comma-seperated {index}:{letter}, where {letter} is not at {index} (0-based) of the target
  -version
        display version
```

### Example

```
$ ./wordler -at 2:i -not-at 1:l,4:k -has l,k -no a,e,o,u
[krill skill skirl]
```


## Web

### Usage

```
Usage of wordler-web:
  -bind string
        bind address (default "0.0.0.0:8080")
  -debug
        enable debug mode
  -max int
        max response length (default 1000)
  -version
        display version
```

### Docker

- [Dockerfile](./Dockerfile)
- [docker-compose.yml](./docker-compose.yml)
