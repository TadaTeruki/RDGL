/*
github.com/TadaTeruki/AkimaSpline/akima_spline.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/


package akima_spline

import(
	"math"
	"sort"
)

type point struct{
	x float64
	y float64
}

type Device struct {
	Point []point
	is_sorted bool
}

func (dev *Device) SetPoint(x, y float64){
	dev.Point = append(dev.Point, point{x, y})
	dev.is_sorted = false
}


func (dev *Device) GetPointByAdress(i int) point {

	if i >= 0 && i < len(dev.Point) {
		return dev.Point[i]
	}

	if i < 0 {
		p3 := dev.GetPointByAdress(0)
		p4 := dev.GetPointByAdress(1)
		p5 := dev.GetPointByAdress(2)
		m4 := (p5.y - p4.y) / (p5.x - p4.x)
		m3 := (p4.y - p3.y) / (p4.x - p3.x)
		x2 := p3.x+p4.x-p5.x
		y2 := p3.y-(p3.x-x2)*(-m4 + 2*m3)

		if i == -1 {
			return point{x2, y2}
		}
		if i == -2 {
			x1 := -p5.x+2*p3.x
			y1 := y2-(x2-x1)*(-2*m4 + 3*m3)
			return point{x1, y1}
		}

	}

	if i >= len(dev.Point){
		ln := len(dev.Point)
		p1 := dev.GetPointByAdress(ln-3)
		p2 := dev.GetPointByAdress(ln-2)
		p3 := dev.GetPointByAdress(ln-1)
		m2 := (p3.y - p2.y) / (p3.x - p2.x)
		m1 := (p2.y - p1.y) / (p2.x - p1.x)

		x4 := -p1.x+p2.x+p3.x 
		y4 := p3.y+(x4-p3.x)*(2*m2-m1)

		if i == ln {
			return point{x4, y4}
		}
		if i == ln+1 {
			x5 := -p1.x+2*p3.x
			y5 := y4+(x5-x4)*(3*m2-2*m1)
			return point{x5, y5}
		}
	}
	

	return point{0, 0}
	
}

func (dev *Device) GetValue(x_tar float64) float64{
	if dev.is_sorted == false {
		sort.Slice(dev.Point, func(i,j int) bool{
			return dev.Point[i].x < dev.Point[j].x
		})
		dev.is_sorted = true
	}

	i2 := sort.Search(len(dev.Point), func(i int) bool { return dev.Point[i].x >= x_tar })
	i2--

	if i2 >= len(dev.Point) {
		i2 = len(dev.Point)-1
	}
	
	p := make(map[int]point, 6)
	m := make(map[int]float64, 6)
	c := 0
	for i := i2-2; i<=i2+3; i++ {
		p[c] = dev.GetPointByAdress(i)
		c++
	}

	for i := 0; i < 6; i++ {
		m[i] = (p[i+1].y - p[i].y) / (p[i+1].x - p[i].x)
	}

	var abs = math.Abs

	w1a := (abs(m[1]-m[0])+abs(m[1]+m[0])*0.5)
	w2a := (abs(m[2]-m[1])+abs(m[2]+m[1])*0.5)
	w1b := (abs(m[3]-m[2])+abs(m[3]+m[2])*0.5)
	w2b := (abs(m[4]-m[3])+abs(m[4]+m[3])*0.5)

	var w1, w2 float64

	if w1a == 0.0 {
		w1 = 0
	} else {
		w1 = w1a/w1b
	}
	
	if w2a == 0.0 {
		w2 = 0.0
	} else {
		w2 = w2a/w2b
	}


	q1 := (m[1]-m[2])/(1.0+w1)
	q2 := (m[3]-m[2])/(1.0+w2)

	a0 := p[2].y
	a1 := q1 + m[2]
	a2 := -(2*q1+q2)/(p[3].x-p[2].x)
	a3 := (q1+q2)/((p[3].x-p[2].x)*(p[3].x-p[2].x))
  
	dx := (x_tar-p[2].x)
	return a0 + a1*dx + a2*dx*dx + a3*dx*dx*dx
}