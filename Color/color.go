package color

type Color struct{
	R, G, B float64
}

type ColorElevation struct {
	color Color
	elevation float64
}

func MakeColor(r, g, b float64) Color{
	var target Color
	target.R = r
	target.G = g
	target.B = b
	return target
}

func MakeColorElevation(r, g, b, elevation float64) ColorElevation{
	var target ColorElevation
	var tc Color
	tc.R = r
	tc.G = g
	tc.B = b
	target.color = tc
	target.elevation = elevation
	return target
}

var elevationList = []ColorElevation{
	MakeColorElevation(0.0, 0.3, 0.7, -10000),
	MakeColorElevation(0.2, 0.7, 0.9, -5000),
	MakeColorElevation(1.0, 1.0, 1.0, -1),
	MakeColorElevation(0.7, 0.95, 0.3, 0),
	MakeColorElevation(0.85, 0.92, 0.4, 300),
	MakeColorElevation(0.95, 0.9, 0.5, 800),
	MakeColorElevation(0.85, 0.75, 0.4, 2000),
	MakeColorElevation(0.75, 0.55, 0.2, 4000),
	MakeColorElevation(0.6, 0.4, 0.1, 10000),
}

func GetColorFromElevation(elevation float64) Color{
	ch := elevationList
	var adove_ch, below_ch ColorElevation
	for i := 0; i<len(ch); i++{
		if i == 0 || ch[i].elevation < elevation{
			below_ch = ch[i]
		} else {
			adove_ch = ch[i]
			break
		}
	}
	var target Color
	var adove_prop float64
	if adove_ch.elevation-below_ch.elevation == 0 {
		adove_prop = 0.0
	} else {
		adove_prop = 1.0-(adove_ch.elevation-elevation)/(adove_ch.elevation-below_ch.elevation)
	}
	
	target.R = adove_ch.color.R*adove_prop + below_ch.color.R*(1.0-adove_prop)
	target.G = adove_ch.color.G*adove_prop + below_ch.color.G*(1.0-adove_prop)
	target.B = adove_ch.color.B*adove_prop + below_ch.color.B*(1.0-adove_prop)
	return target
}