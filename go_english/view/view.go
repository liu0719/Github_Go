package view

import (
	"english/control"
	"english/model"
	"english/utils"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type View struct {
	WordBox        *canvas.Text    //单词的区域
	Wordlist       *model.WordList //获取的单词
	entry          *widget.Entry   //输入框
	i              int
	w              fyne.Window
	text           *canvas.Text
	TotalErrNumber int
	NewWordlist    []*model.Word
}

func (v *View) Begin() *widget.Box {
	iconLaba, err := utils.ReadIcon("./asset/laba.svg")
	if err != nil {
		log.Println("图片读取出错", err)
	}
	//发音操作
	button2 := widget.NewButtonWithIcon("发音", fyne.NewStaticResource("laba", iconLaba), func() {
		proCon := control.PronounceControl{}
		go proCon.OverAllControl(v.Wordlist.Words[v.i])
	})
	button3 := widget.NewButton("确认", func() {
		defer func() {
			if err := recover(); err != nil {
				err = fmt.Errorf("出错在: %v", err)
			}
		}()
		v.IsTure()
		v.i = rand.Intn(len(v.Wordlist.Words))
		v.NextWord()
	})
	button3.Resize(fyne.NewSize(100, 30))
	// 创建输入框并设置大小
	v.entry = widget.NewEntry()
	v.entry.Resize(fyne.NewSize(200, 70))
	// 将输入框添加到容器中
	box := container.NewVBox()
	box.Add(v.entry)
	//提示,用canvas制作并输出
	v.text = canvas.NewText("请输入完成后点击确认", color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	})
	//居中
	v.text.Alignment = fyne.TextAlignCenter
	v.text.TextSize = 30
	v.i = rand.Intn(len(v.Wordlist.Words))
	v.WordBox = canvas.NewText(v.Wordlist.Words[v.i].Chinese, color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	})
	v.text.TextSize = 30
	v.text.Alignment = fyne.TextAlignCenter

	content := widget.NewVBox(
		v.WordBox,
		widget.NewHBox(button2),
		box,
		button3,
		v.text,
	)
	return content
}

// 图形界面初始化
func (v *View) ViewInit() {
	os.Setenv("FYNE_FONT", "./asset/fonts.ttc") //设置env环境

	//创建一个窗口
	a := app.New()
	//窗口标题
	v.w = a.NewWindow("单词背诵")
	//窗口图标
	ico, err := utils.ReadIcon("./asset/game.png")
	if err != nil {
		log.Println("图片读取出错", err)
	}
	v.w.SetIcon(fyne.NewStaticResource("ico", ico))
	//窗口大小
	v.w.Resize(fyne.NewSize(1100, 500))
	// 窗口居中
	v.w.CenterOnScreen()
	button1 := widget.NewButton("开始!", func() {
		content := v.Begin()
		v.w.SetContent(content)
	})
	content := container.NewCenter(button1)
	v.w.SetContent(content)

	//展示窗口
	v.w.ShowAndRun()
}

// 单词下一个
func (v *View) NextWord() {
	v.WordBox.Text = v.Wordlist.Words[v.i].Chinese
	v.WordBox.Refresh()
}

// 单词判断对错
func (v *View) IsTure() {
	flag := false
	if v.entry.Text == v.Wordlist.Words[v.i].English {
		flag = true
		v.NewWordlist = append(v.NewWordlist, v.Wordlist.Words[v.i])
		v.Wordlist.Words = append(v.Wordlist.Words[:v.i], v.Wordlist.Words[v.i+1:]...)
	}

	v.entry.Text = ""
	v.entry.Refresh()

	str := ""
	if flag {
		str = "正确"
		v.text.Color = color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		}
	} else {
		str = "错误,[" + v.Wordlist.Words[v.i].Chinese + "]的英语为:[" + v.Wordlist.Words[v.i].English + "]"
		v.text.Color = color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		}
		v.Wordlist.Words[v.i].ErrNumber++
		v.TotalErrNumber++
	}
	v.text.Text = str
	v.text.Refresh()
	if len(v.Wordlist.Words) == 0 {

		if v.TotalErrNumber > 0 {
			v.Total()
		} else {
			v.IsFinal()
		}
	}
	go func() {
		time.Sleep(10 * time.Second)
		v.text.Text = "请输入完成后点击确认"
		v.text.Color = color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		}
		v.text.Refresh()
	}()

}
func (v *View) IsFinal() {
	if len(v.Wordlist.Words) == 0 {
		bye := canvas.NewText("已经背完所有单词辣,休息一下吧", color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		})
		//居中
		bye.Alignment = fyne.TextAlignCenter
		bye.TextSize = 30
		v.w.SetContent(bye)
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
}
func (v View) Total() {
	content := widget.NewVBox()
	label1 := widget.NewLabel("错误情况:")
	content.Children = append(content.Children, label1)
	fmt.Println(v.TotalErrNumber)
	for _, value := range v.NewWordlist {
		if value.ErrNumber != 0 {
			str := "[" + value.English + "]:[" + value.Chinese + "]," + "错" + fmt.Sprintf("%v", value.ErrNumber) + "次"
			label := widget.NewLabel(str)
			content.Children = append(content.Children, label)
		}

	}
	button4 := widget.NewButton("确定", func() {
		v.IsFinal()
	})
	content.Children = append(content.Children, button4)
	v.w.SetContent(content)

}
