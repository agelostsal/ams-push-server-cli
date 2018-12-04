# Ams-push-server-cli

An command line tool to interface with the [ams push server](https://github.com/ARGOeu/ams-push-server).

# How to use
**Note:** The `-u/--uri` flag is mandatory across all commands.
```commandline
Usage:
   [command]

Available Commands:
  activate    Activates a subscription on a push server
  deactivate  Deactivates a subscription on a push server
  health      Performs a health check call
  help        Help about any command

Flags:
  -h, --help         help for this command
  -u, --uri string   -u host:port

Use " [command] --help" for more information about a command.
```

### Perform a health check on a push server
```commandline
$ apscli health -u example.com:443
```
```commandline
$ Success: SERVING
```

### Inform a push server to activate the push functionality for the given subscription

```commandline
The activate command informs a push server to start the push functionality for the given subscription

Usage:
   activate

Flags:
  -s, --full-sub string        -s /projects/projectname/subscriptions/subanme
  -f, --full-topic string      -f /projects/projectname/topics/topicname
  -h, --help                   help for activate
  -e, --push-endpoint string   -e https://127.0.0.1:5000/receive_here
  -p, --retry-period uint32    -p 300 (default 300)
  -t, --retry-type string      -t linear (default "linear")

Global Flags:
  -u, --uri string   -u host:port
```



```commandline
$ apscli activate -u example.com:443 -s /projects/projectname/subscriptions/subanme -f /projects/projectname/topics/topicname -e https://example.net:443/receive_here -p 300 -t linear
```
```commandline
$ Success: Subscription /projects/projectname/subscriptions/subanme activated
```


### Inform a push server to deactivate the push functionality for the given subscription

```commandline
The deactivate command informs a push server to start the push functionality for the given subscription

Usage:
   deactivate

Flags:
  -s, --full-sub string        -s /projects/projectname/subscriptions/subanme
  -h, --help                   help for deactivate

Global Flags:
  -u, --uri string   -u host:port
```

```commandline
$ apscli deactivate -u example.com:443 -s /projects/projectname/subscriptions/subanme 
```
```commandline
$ Success: Subscription /projects/projectname/subscriptions/subanme deactivated
```