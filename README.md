# Dwnldr

This is a tool to help automate the downloading of files. It does not download anything itself, but acts as a router to call the right programs based on the parameter it has been passed, which is expected to be a URL.

You will need to install the tools in the config file yourself.

## Usage

You will need a configuration file first. There is an example one in `config/dwnldr.toml`.

```
dwnldr download "https://www.youtube.com/watch?v=XfELJU1mRMg"
```
