A little command line tool to track how much time you spend on different tasks.

## Usage

General workflow:

```
# Add a new task to your task list
$ go run main.go add-task
```

Run `go run main.go --help` for full usage.

## Development

This CLI uses [cobra](https://github.com/spf13/cobra). See documentation for more information. Install the cobra CLI to auto-generate code for new commands.

```
# Add a new command
$ cobra add commandName
```
