### Building

#####Windows
cmd:
```
cd <PROJECT-ROOT>
go build -v -o bin\tinyfilter.exe dev\tinyfilter.go
```
git bash:
```
cd <PROJECT-ROOT>
go build -v -o bin/tinyfilter.exe dev/tinyfilter.go
```

#####FreeBSD
```
cd <PROJECT-ROOT>
go build -v -o bin/tinyfilter dev/tinyfilter.go
```


### Testing

##### Remote

1. Install go, dlv
    - Golang, [Getting started](https://golang.org/doc/install)
    - Delve, [installation](https://github.com/go-delve/delve/tree/master/Documentation/installation)

2. Setup remote debug in IDEA, see JetBrains blog for [instructions](https://blog.jetbrains.com/go/2019/02/debugging-with-goland-getting-started/#debugging-a-running-application-on-a-remote-machine).
    - Host: ip of remote computer 
    - port: 2345
    
3. Remote, build in terminal
    ```bash
    cd <PROJECT-ROOT>
    go build -v -o bin/tinyfilter dev/tinyfilter.go
    ```

4. Remote, start debugger in terminal
    ```bash
    dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./bin/tinyfilter r
    ```

5. Run debug config. 
    You can run it multiple time without re-compilation.

6. Kill Delve, other shell window
    ```bash
     killall -HUP dlv
    ```

Other option is running tests at remote computer. 
For do that, skip step 3 and use other command at step 4:
```bash
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient test ./dev/util/exec/
```
Please note need show folder containing test code. 