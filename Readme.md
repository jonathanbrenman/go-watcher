# Watch directory and subdirectories files
## Watcher files made with golang

### Start watcher
```
$ go run main.go watch --path ./prueba --debug true --filters mp4,txt --delay 5s
```

### List files recursive once
```
$ go run main.go list --path ./prueba --filters txt,mp4
```