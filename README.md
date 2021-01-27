# tapioca_logging

tapioca_logging is a simple GO module to help with logging to a file.  It uses build in go packages (fmt, log, os) to help you add logging to your projects quickly.

## Installation
tapioca_logging can be added directly to your go projects via import:

```Go
include (
    "github.com/TapiocaTechnologies/tapioca_logging"
)
```
Make sure you are using the most current version [v0.1.3](https://github.com/TapiocaTechnologies/tapioca_logging/releases/tag/v0.1.2) you can change the version you are using in your go.mod file.


## Function / Usage

**SetLogFile(path string):** Allows you to set where you would like to save your log file, along with what you would like it named.  This function allows relative and absolute paths

```
tapioca_logging.SetLogFile("testlogfile.log")
```
```
tapioca_logging.SetLogFile("c:\dev\testlogfile.log")
```

**Logging(message string):** This is a general function that allows you to write a string to a file.

```
message := "Welcome to tapioca_logging"
tapioca_logging.Logging(message)
```

**LoggingError(err error):** This is a general function that allows you to write an error to a file.

```
errs := errors.New("Error")
tapioca_logging.LoggingError(errs)
```

### Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)



