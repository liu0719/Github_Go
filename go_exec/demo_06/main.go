package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	width     = 400 // 图片宽度
	height    = 400 // 图片高度
	margin    = 20  // 边距
	axisColor = 0   // 坐标轴颜色
	lineColor = 0   // 约束条件线条颜色
)

// 绘制函数
func draw() image.Image {
	// 创建灰色背景
	img := image.NewGray(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetGray(x, y, color.Gray{Y: 255})
		}
	}

	// 绘制坐标系
	for x := margin; x <= width-margin; x++ {
		img.SetGray(x, height-margin, color.Gray{Y: axisColor})
	}
	for y := margin; y <= height-margin; y++ {
		img.SetGray(margin, y, color.Gray{Y: axisColor})
	}

	// 计算约束条件
	// 4x1 + 3x2 <= 120 → x2 <= (-4/3)x1 + 40
	// x1 + 2x2 <= 60 → x2 <= (-1/2)x1 + 30
	// x1 >= 0, x2 >= 0
	xMax := 60.0
	yMax := 40.0
	x1 := 0.0
	//x2 := 0.0

	// 绘制约束条件
	for x := x1; x <= xMax; x += 0.1 {
		y := (-4.0/3.0)*x + yMax
		if y <= yMax && y >= 0 {
			img.Set(int(x)+margin, height-margin-int(y), color.Gray{Y: lineColor})
		}
	}
	for x := x1; x <= xMax; x += 0.1 {
		y := (-1.0/2.0)*x + 30.0
		if y <= yMax && y >= 0 {
			img.Set(int(x)+margin, height-margin-int(y), color.Gray{Y: lineColor})
		}
	}

	// 返回图片
	return img
}

func main() {
	// 创建图片文件
	f, err := os.Create("fruit.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 绘制并保存图片
	png.Encode(f, draw())
}
