package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"time"
)

func SaveFile(file *multipart.FileHeader) (filepath, dst string, err error) {

	// 2、获取后缀名判断类型是否正确.jpg .png .gif .jpeg
	//path.Ext()可以获取文件后缀名
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", "", errors.New("文件类型不合法")
	}
	// 3、创建图片保存目录static/upload /20200623
	day := time.Now().Format("20060102")
	dir := "views/static/img/upload/" + day
	//该方法可以创建多层目录
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		return "", "", errors.New("创建文件夹失败")
	}
	// 4、生成文件名称和文件保存的目录
	unix := time.Now().Unix()
	filename := fmt.Sprintf("%v%v", unix, extName)

	//5、执行上传
	//路径拼接
	dst = path.Join(dir, filename)
	filepath = path.Join("/static/img/upload/", day, filename)
	return
}
