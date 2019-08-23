# goenvdir
This utility is a wrapper around the common execution of any programme, except
it needs to get environment variables from the specified folder first and then
run the programme passed to it as an argument.

See: https://cr.yp.to/daemontools/envdir.html for reference.

## Usage

```goenvdir [FLAGS] /path/to/env/dir some_prog```

## Supported flags

### clean (-c)
Clears out all the environment before using the passed ones.

## TODO
1) Add the ability to read files in various formats (JSON, YAML, INI etc.)