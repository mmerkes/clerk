A little command line tool to track how much time you spend on different tasks.

## Usage

General workflow:

```
# Add a new task to your task list
$ go run main.go add
# Start a task of ID 1
$ go run main.go start -i 1
# Stop a task of ID 1
$ go run main.go stop -i 1
```

Run `go run main.go --help` for full usage.

# Install on your machine

You can install clerk on your machine to run the commands without having to build the binary every time.

```
# Install the binary to your GOPATH
$ go build -o $GOPATH/bin/clerk
# Call clerk
$ clerk list
```

NOTE: The binary you installed will use the same `~/.clerk-db` that is used in development until that is configured for the dev environment.

## Development

This CLI uses [cobra](https://github.com/spf13/cobra). See documentation for more information. Install the cobra CLI to auto-generate code for new commands.

```
# Add a new command
$ cobra add commandName
```

Format code:

```
$ ./gofmt
```
