

![result-bg2](https://user-images.githubusercontent.com/57752033/144701904-1a87e028-6904-4f99-93d4-062fac35c45b.png)

# RDGL - Realistic procedural DEM generation library

by TadaTeruki (Perukii) <tadateruki.public@gmail.com> <br>
written in **Golang**

## Features

RDGL generates various procedural-terrains like :
 - Plain
 - Valley
 - Rias coast
 - Continental shelf
 - Mountain range
<br>and so on ...


### Procedural DEM generation

|Seed = 8|Seed=14<br>LandProportion01=0.85|Seed = 3<br>LandProportion01=0.15|Seed = 25|
|---|---|---|---|
|<img src="https://user-images.githubusercontent.com/57752033/144703940-64409bae-8279-4bf1-9b6c-df8f3a44ce2d.png" height="200px">|<img src="https://user-images.githubusercontent.com/57752033/144704213-c1d2c452-8970-40fb-a219-f904dbf18d1b.png" height="200px">|<img src="https://user-images.githubusercontent.com/57752033/144704154-9f09cbc2-91d1-4ee1-91bb-d8e977acddf8.png" height="200px">|<img src="https://user-images.githubusercontent.com/57752033/144704341-eaa2d1ca-49d2-4889-8847-26b7fce88692.png" height="200px">|
```go
package main
import output "github.com/TadaTeruki/RDGL/Output"
import rdg "github.com/TadaTeruki/RDGL"

func main(){
  rdg.EnableProcessLog()
  var seed int64 = 14
  dem := rdg.NewDEM(seed)
  dem.Generate()
  output.WriteDEMtoPNGwithShadow("result.png", &dem, 300, -1, output.DefaultShadow(&dem))
  // details : examples/hello_dem.go, examples/hello_dem_detailed.go, examples/write_to_png_with_shadow.go
}
```

___

### Outline interpolation (Preparing)

|examples/resources/draft.png|examples/resources/swan.png|
|---|---|
|<img src="https://user-images.githubusercontent.com/57752033/144703651-cc438a8d-84e3-4ac7-bd37-e10074ad2340.png" height="150px"><br><img src="https://user-images.githubusercontent.com/57752033/144703715-acad18ba-f2c9-4438-aac4-712b112b80e6.png" height="150px">|<img src="https://user-images.githubusercontent.com/57752033/144702040-b51fb5fa-a7f5-4cfb-9bd8-4950b1d05734.jpg" height="150px"><br><img src="https://user-images.githubusercontent.com/57752033/144703435-9a51b668-8640-4ac8-aa9c-0f36871f224d.png" height="150px">|

___

### TXT/PNG/OBJ output

|PNG|OBJ (3D model) [*1] [*2]|
|---|---|
|<img src="https://user-images.githubusercontent.com/57752033/144703530-7a11bd6b-ef2f-4f66-bf7f-e2b42098eedc.png" height="150px">|<img src="https://user-images.githubusercontent.com/57752033/144702174-8a3e0c2b-1645-4f2e-a991-e5ac7ea8e615.gif" height="150px">|

[*1] displayed with https://github.com/RBFraphael/meshviewer <br>
[*2] Elevation = x50

```go
...
func main(){
  ...
  output.WriteDEMtoTXT("result.txt", &dem, 1000, -1)
  output.WriteDEMtoPNG("result.png", &dem, 300, -1)
  output.WriteDEMtoPNGwithShadow("result1.png", &dem, 300, -1, output.DefaultShadow(&dem))
  output.WriteDEMtoOBJ("result.obj", &dem, 100, -1, 5.0, false)
  // details : examples/write_to_txt.go, examples/write_to_png.go,
  //           examples/write_to_png_with_shadow.go, examples/write_to_obj.go
}
```

___

## Installation

```
$ go get github.com/TadaTeruki/RDGL
```

## Dependencies

[![go-cairo](https://github-readme-stats.vercel.app/api/pin/?username=ungerik&repo=go-cairo)](https://github.com/ungerik/go-cairo)<br>
[![gods](https://github-readme-stats.vercel.app/api/pin/?username=emirpasic&repo=gods)](https://github.com/emirpasic/gods)<br>
[![gospline](https://github-readme-stats.vercel.app/api/pin/?username=cnkei&repo=gospline)](https://github.com/cnkei/gospline)<br>

## LICENSE

LGPL - v3.0
