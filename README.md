# tapioca_logging

tapioca_logging is a simple GO module to help with logging to a file.  It uses build in go packages (fmt, log, os) to help you add logging to your projects quickly.

## Installation
tapioca_logging can be added directly to your go projects via import:

```Go
import (
    "github.com/TapiocaTechnologies/tapioca_logging"
)
```
Make sure you are using the most current version [v0.1.4](https://github.com/TapiocaTechnologies/tapioca_logging/releases/tag/v0.1.4) you can change the version you are using in your go.mod file.


## Function / Usage

**SetLogFile(path string):** Allows you to set where you would like to save your log file, along with what you would like it named.  This function allows relative and absolute paths

```
tapioca_logging.SetLogFile("testlogfile.log")
```
```
tapioca_logging.SetLogFile("c:\dev\testlogfile.log")
```

**Logging(message string):** This is a general function that allows you to write a string to a file. Inserts log message with [GENERAL] tag before message.

```
message := "Welcome to tapioca_logging"
tapioca_logging.Logging(message)
```

**LoggingError(err error):** This is a general function that allows you to write an error to a file. Inserts log message with [ERROR] tag before message.

```
errs := errors.New("Error")
tapioca_logging.LoggingError(errs)
```

**LoggingNetwork(message string):**  Use this function if you need to log anything that has to do with networking. Inserts log message with [NETWORK] tag before message.

```
message := "Client successfully connected"
tapioca_logging.LoggingNetwork(message)
```

**LoggingNetwork(err error):**  Use this function if you need to log anything that has to do with networking errors. Inserts log message with [NETWORK] [ERROR] tag before message.

```
errs := errors.New("Client failed to connect")
tapioca_logging.LoggingNetworkError(errs)
```

## Logfile output

```
2021/01/26 19:01:46 [GENERAL] Welcome to tapioca_logging
2021/01/26 19:01:46 [ERROR] Error
2021/01/26 19:01:46 [NETWORK] Client successfully connected
2021/01/26 19:01:46 [NETWORK] [Error] Client failed to connect
```


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)



