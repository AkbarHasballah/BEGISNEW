# atapi

An Iteung API Interface HTTTP Method

```go
package main

import (
 "fmt"
 "github.com/aiteung/atapi"
)

type TestApi struct {
 Phone      string `json:"phoneNumber"`
 Password   string `json:"password"`
 FirebaseId string `json:"firebaseId"`
 DeviceId   string `json:"deviceId"`
}

type Sister struct {
 Id_sdm string `url:"id_sdm" json:"id_sdm"`
}

type Response struct{
 Message      string `json:"message"`
 Status   string `json:"status"`
}

func main(){
 var res Response
 dt := TestApi{
    Phone:      "+6285155476774",
    Password:   "#P@ssw0rd",
    FirebaseId: "123",
    DeviceId:   "6580fb6e714844ca",
 }
 url := "https://awangga.requestcatcher.com/"
 res = atapi.PostStruct[Response](dt, url)
 fmt.Println("TestPostStruct : ", res)
 res = atapi.Get[Response](url)
 fmt.Println("TestPostStruct : ", res)
 res = atapi.GetStructWithToken[Response]("token", "dsfdsfdsfdsfdsf", dt, url)
 fmt.Println("GetStructWithToken : ", res)
 res = PostStructWithToken[Response]("Login", "dsfdsfdsfdsfdsf", dt, url)
 fmt.Println("PostStructWithToken : ", res)
 res = PostStructWithBearer[Response]("dsfdsfdsfdsfdsf", dt, url)
 fmt.Println("PostStructWithBearer : ", res)
 res = GetStructWithBearer[Response]("dsfdsfdsfdsfdsf", dt, url)
 fmt.Println("GetStructWithBeare : ", res)
 res = GetStruct[Response](dt, url)
 fmt.Println("GetStruct : ", res)

}

```

## Tagging

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/aiteung/atapi@v0.0.1
```
