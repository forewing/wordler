# WORDLER

## Usage

```
./wordler -h
Usage of ./wordler:
  -at {index}:{letter}
        comma-seperated {index}:{letter}, where {letter} is at {index} (0-based) of the target
  -has string
        comma-seperated strings the target contains
  -len int
        length of word (default 5)
  -no string
        comma-seperated strings the target does not contain
  -not-at {index}:{letter}
        comma-seperated {index}:{letter}, where {letter} is not at {index} (0-based) of the target
```

## Example

```
$ ./wordler -at 2:i -not-at 1:l,4:k -has l,k -no a,e,o,u
[krill skill skirl]
```
