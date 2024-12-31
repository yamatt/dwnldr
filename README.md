# Dwnldr

This is a tool to help automate the downloading of files. It does not download anything itself, but acts as a router to call the right programs based on the parameter it has been passed, which is expected to be a URL.

You will need to install the tools in the config file yourself.

This is my first serious attempt at a Go program. So it's a bit of a throw-away project. You can do this program with about 20 lines of bash.

## Usage

### Download

You will need a configuration file first. There is an example one in `config/dwnldr.toml`.

```bash
dwnldr download "https://www.youtube.com/watch?v=XfELJU1mRMg"
```

### Configure

The `dwnldr.toml` file lists all the types of handlers. You can add new ones or remove existing ones.

Each handler must have:

- **name**: A human readable reference for what the type of handler is
- **pattern** A regular expression that matches for the URL that is passed in to the program
- **command** An array of shell arguments that are run when the pattern is matched. You can also populate arguments with `{url}` and `{outputDir}` to provide additional information to the program.

Note that the order of the handlers matters, and it will only run one handler. If multiple handlers match then only the first one will run. Ones at the top will be tested to match the pattern first. If you put the HTTP handler near the top it will likely grab everything without passing it off to other programs.

**Example:**

```toml
[[handlers]]
name = "Generic HTTP Downloader"
pattern = 'https?://'
command = ["curl", "{url}", "--output-dir", "{outputDir}"]
```

## Compile

```bash
go build -o dwnldr ./cmd/dwnldr
```

## Install

```bash
go install github.com/yamatt/dwnldr/cmd/
```
