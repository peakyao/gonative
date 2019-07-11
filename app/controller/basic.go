package controller

import (
	"fmt"
	"gonative/libs/slog"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
)

func Study(w http.ResponseWriter, r *http.Request) {
	slog.InfoLog("Basic Study")

	var str string = "abcde"
	// stringFor(str)
	// stringRange(str)

	n := checkBool(true)
	slog.InfoLog("checkBool result:", n)

	// signImage()

	str = `第一行
		   第二行
		   第三行
	`

	w.Write([]byte(str + "中国人"))
}

//字符串循环
func stringFor(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])

		fmt.Printf("%c", str[i])
		fmt.Println()
	}
}

func stringRange(str string) {
	for key, value := range str {
		fmt.Println(key, value)
		fmt.Printf("%c %c", key, value)
	}
}

//输出sign函数图像
func signImage() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}

func checkBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
