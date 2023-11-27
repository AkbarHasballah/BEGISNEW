# dzikri7
Dzikri


```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v1.0.0                          #set tag version
git push origin v1.0.0                  #push tag version to repo
GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
go get example.com/mymodule@v0.1.0
go list -m example.com/mymodule@v0.1.0   #publish to pkg dev, replace ORG/URL with your repo URL
