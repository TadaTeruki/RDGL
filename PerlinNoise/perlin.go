package perlin

import(
	"math"
	"math/rand"
)

// ノイズを提供するオブジェクト
type PerlinObject struct{
	Seed []int
}

// ノイズのシード値を設定
func (obj* PerlinObject) SetSeed(seed_num int64){
	var p []int
	for i:=0; i<256; i++{
		p = append(p, i)
	}
	rand.Seed(seed_num)
	rand.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })

	for i:=0; i<2; i++{
		for j:=0; j<256; j++{
			obj.Seed = append(obj.Seed, p[j])
		}
	}
}

// 値を丸める関数
func getFade(t float64) float64{
	return t*t*t*(t*(t*6-15)+10)
}

// 1次関数
func getLerp(a,b,x float64) float64{
	return a+x*(b-a)
}


func makeGrad(hash int,x,y,z float64) float64{
    switch(hash & 0xF){
        case 0x0: return x + y;
        case 0x1: return -x + y;
        case 0x2: return  x - y;
        case 0x3: return -x - y;
        case 0x4: return  x + z;
        case 0x5: return -x + z;
        case 0x6: return  x - z;
        case 0x7: return -x - z;
        case 0x8: return  y + z;
        case 0x9: return -y + z;
        case 0xA: return  y - z;
        case 0xB: return -y - z;
        case 0xC: return  y + x;
        case 0xD: return -y + z;
        case 0xE: return  y - x;
        case 0xF: return -y - z;
        default: return 0;
    }
}

func getGrad(hash int,x,y,z float64) float64{
	return makeGrad(hash&15,x,y,z)
}

// ノイズを生成する関数
func (obj* PerlinObject) SetNoise(x,y,z float64) float64{

	var xi, yi, zi int
	var xf, yf, zf float64
	var u, v, w float64

	xi = int(math.Floor(x))%255
	yi = int(math.Floor(y))%255
	zi = int(math.Floor(z))%255

	xf = x - math.Floor(x)
	yf = y - math.Floor(y)
	zf = z - math.Floor(z)

	u = getFade(xf)
    v = getFade(yf)
    w = getFade(zf)

    var aaa, aba, aab, abb, baa, bba, bab, bbb int
	p := obj.Seed
    aaa = p[p[p[xi  ]+yi ] +zi  ]
    aba = p[p[p[xi  ]+yi+1]+zi  ]
    aab = p[p[p[xi  ]+yi  ]+zi+1]
    abb = p[p[p[xi  ]+yi+1]+zi+1]
    baa = p[p[p[xi+1]+yi  ]+zi  ]
    bba = p[p[p[xi+1]+yi+1]+zi  ]
    bab = p[p[p[xi+1]+yi  ]+zi+1]
    bbb = p[p[p[xi+1]+yi+1]+zi+1]

	var x1,x2,y1,y2 float64
	x1 = getLerp(getGrad(aaa,xf,yf,zf), getGrad(baa,xf-1,yf,zf),u)
	x2 = getLerp(getGrad(aba,xf,yf-1,zf), getGrad(bba,xf-1,yf-1,zf),u)
	y1 = getLerp(x1, x2, v)

	x1 = getLerp(getGrad(aab,xf,yf,zf-1), getGrad(bab,xf-1,yf,zf-1),u)
	x2 = getLerp(getGrad(abb,xf,yf-1,zf-1), getGrad(bbb,xf-1,yf-1,zf-1),u)
	y2 = getLerp(x1, x2, v)

	return getLerp(y1, y2, w)

}

// ノイズを出力する関数
func (obj* PerlinObject) Noise(x,y,z float64) float64{
    return obj.SetNoise(x, y, z)*0.5+0.5
}

// オクターブ付きのノイズを生成する関数 (octaves:細分処理の回数, persistence:地形のきめ細やかさ)
func (obj* PerlinObject) SetOctaveNoise(octaves int, persistence,x,y,z float64) float64 {
    value := 0.0
    amp := 1.0
    for i := 0; i < octaves; i++{
        value += obj.SetNoise(x,y,z) * amp
        x *= 2.0
        y *= 2.0
        z *= 2.0
        amp *= persistence
    }
    return value
}

// オクターブ付きのノイズを出力する関数
func (obj* PerlinObject) OctaveNoise(octaves int, persistence,x,y,z float64) float64{
    return obj.SetOctaveNoise(octaves, persistence, x, y, z)*0.5+0.5
}