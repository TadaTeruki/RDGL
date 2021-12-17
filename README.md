

![result-bg2](https://user-images.githubusercontent.com/57752033/144701904-1a87e028-6904-4f99-93d4-062fac35c45b.png)

# RDGL - Realistic procedural DEM generation library

**DEM** : Digital Elevation Model

 - [**Online Trial**](https://go.dev/play/p/g4kX79ADAoY) for Go Playground is now available!

## Features

RDGL can generate procedural terrains, including...
 - Plain
 - Valley
 - Rias coast
 - Continental shelf
 - Mountain range
 <br>... and so on.


### Random DEM generation

|Seed = 8|Seed=14<br>LandProportion01=1.0|
|---|---|
|<img src="https://user-images.githubusercontent.com/69315285/146559610-53b8e21f-3574-4cff-a7ee-0b7c13c32c13.png" width="200px" height="200px">|<img src="https://user-images.githubusercontent.com/69315285/146560007-9976a0e6-a81e-4533-a7aa-e200611d8a06.png" width="200px" height="200px">|

|Seed = 14<br>LandProportion01=0.15|Seed = 2022|
|---|---|
|<img src="https://user-images.githubusercontent.com/69315285/146560427-846d42d7-1350-4d08-bdcf-3ec0dff7d839.png" width="200px" height="200px">|<img src="https://user-images.githubusercontent.com/69315285/146560806-84541b47-66ef-4229-a95b-e2ce73d6c1b1.png" width="200px" height="200px">|

[*] 1000x1000 (km2) terrain

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

## Author & Contributors

Author : Tada Teruki < tadateruki.public@gmail.com >

Copyright (c) 2021 Tada Teruki

![result](https://user-images.githubusercontent.com/69315285/146564971-b1a510b2-d5c9-4ca2-91cd-277ffbb6d1c7.png)
