# Module Handling Iteung
Iteung module routing

Usage :
declare : WAIface model.IteungWhatsMeowConfig, DBIface model.IteungDBConfig
```go
Modulename,Pesan:=IteungModuleCall(WAIface, DBIface)
if Modulename != ""{
    resp, err := CallAndSend(Modulename, Pesan, WAIface)
    if err != nil {
				fmt.Println("Gagal Kirim Whatsapp Dari ITEUNG V2 Baru")
				fmt.Println(err)
			}
			fmt.Println(resp)
}
```
Status Message  
![image](https://github.com/aiteung/module/assets/11188109/fc5e3fbf-a32d-4bd8-8b91-e6dd90dcba98)


List of Mongo Collection :

1. Typo : to replace typo word
2. Module : return module name from keyword 

## Dev
```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.20                         #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/aiteung/module@v0.0.20   #publish to pkg dev, replace ORG/URL with your repo URL
```
