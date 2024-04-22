# timenow

A cross-platform CLI tool to get the current time in different timezones and formats.

## Install

```bash
go install github.com/josestg/timenow@latest    
```

## Usage

```bash
# by default, it will show the current time in unix epoch format (seconds).
timenow

# to show the current time in a unix epoch format (milliseconds).
timenow -format=epochs-millis

# to show the current time in a RFC3339 format and Asia/Jakarta timezone.
timenow -format=rfc3339 -timezone=Asia/Jakarta
```
