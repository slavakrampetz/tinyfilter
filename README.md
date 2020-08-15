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

### Automate using Task (TBF)

##### Task, links: 
1. [Task](https://taskfile.dev/)
2. [Styleguide](https://taskfile.dev/#/styleguide)
3. [Example web app](https://github.com/go-task/examples)
    - Direct link to [Taskfile.yml](https://github.com/go-task/examples/blob/master/go-web-app/Taskfile.yml)


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
Please note need set folder containing test code. 


### Links

1. Echo, https://echo.labstack.com/
2. GoByExample, https://gobyexample.com/
3. WEB, slava: http://192.168.1.21:8085/

### Commands
1. Ping: test service online
    * http://HOST:PORT/c/ping/
2. Reload tinyproxy: restart service
    * http://HOST:PORT/c/reload/?key=KEY
3. Get a current status of tinyproxy
    * http://HOST:PORT/c/youtube/get/?key=KEY
4. Turn Youtube on/off: 
   - change config of tinyproxy
   - restart tinyproxy service
   - commands
     http://HO.ST:PORT/c/youtube/on/?key=KEY -- youtube on
     http://HO.ST:PORT/c/youtube/off/?key=KEY -- youtube off
### Commands, how to run


