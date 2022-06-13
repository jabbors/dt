# dt

dt is a command line tool to aid transforming durations (typically Golang printed durations) into a normalized format that can be used for further processing.

## Installation

```
go get github.com/jabbors/dt
```

## Example use case

Say you have to following input in a file named `sample.dat` and would like to know the `max`, `min` and `median` value

```
3.765963ms
1.782821ms
2.35635ms
3.14281545s
47.767401ms
3.566412ms
2.012246ms
35.296362ms
2.46694ms
3.153943706s
```

Then you execute

```
$ cat sample.dat | dt | datamash max 1 min 1 median 1
3153.943706 1.782821 3.6661875
```
