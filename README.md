# Geographic Information System  By Akbar Hasbullah

## 1. Geowithin
Menentukan apakah suatu data lokasi berada di dalam suatu bentuk geometri (polygon) tertentu.

Format 
```
{
  "type":"",
  "coordinates":[],
  "min":,
  "max":
}
```

## 2. GeoIntersect
Menentukan apakah suatu bentuk geometri (polygon) memotong (intersect) dengan data lokasi yang disimpan dalam MongoDB.

Format
```
{
    "type": "LineString/Point/Polygon",
    "coordinates": 
                [[],
                []] 

}
```
## 3.Box
 Menentukan sebuah kotak (box) di sekitar suatu wilayah tertentu dalam sistem koordinat geografis. Data yang memiliki lokasi di dalam kotak tersebut akan dipilih.

 Format
 ```
 {
    "coordinates":[ [ Koordinan Kiri Bawah],
     [ Koordinat Kanan Atas] ]
}
 ```
 ## 4. Near
 Mencari data lokasi yang berada paling dekat dengan suatu titik atau koordinat tertentu.

 Format

 ```
 {
  "coordinates":[],
  "min":
  "max":
}
 ```

 ## 5. NearSphere
 Operator ini mirip dengan $Near, namun dengan tambahan parameter jari-jari (radius) dalam radians.

 Format 

 ```
 {
  "type":"Point",
  "coordinates":[
  ],
  "min":
  "max":
}
 ```

 ## 6. Center
 Menentukan sebuah titik tengah (center) dari suatu wilayah tertentu dalam sistem koordinat geografis. Data yang memiliki lokasi di sekitar titik tengah tersebut akan dipilih.

 Format

 ```
 {
    "coordinates": [] ,
    "radius": 
}
 ```

Rilis Package
```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v1.0.0                          #set tag version
git push origin v1.0.0                  #push tag version to repo
GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
go get example.com/mymodule@v0.1.0
go list -m example.com/mymodule@v0.1.0   #publish to pkg dev, replace ORG/URL with your repo URL
