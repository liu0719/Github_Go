package admin

import (
	"fmt"
	"go_gin/server_02/models"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	/*
		c.String(http.StatusOK, "--用户列表--")
		users := []models.User{}
		models.DB.Find(&users)
		c.JSON(http.StatusOK, gin.H{
			"result": users,
		})
	*/
	//查询年龄大于20的用户
	users := []models.User{}
	//指定查询
	models.DB.Where("age>20").Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}
func (con UserController) Add(c *gin.Context) {
	// c.HTML(http.StatusOK, "useradd.html", "")
	user := models.User{
		Username: "刘广硕",
		Age:      20,
		Email:    "335293667@qq.com",
		AddTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	tx := models.DB.Create(&user)
	fmt.Println(tx, "++++", user)
	c.String(http.StatusOK, "增加数据成功")
}
func (con UserController) AddSame(c *gin.Context) {
	c.HTML(http.StatusOK, "useraddsame.html", "")
}
func (con UserController) Edit(c *gin.Context) {
	// c.HTML(http.StatusOK, "useredit.html", "")

	/*
		user := models.User{Id: 3}
		models.DB.Find(&user)
		//更新数据
		user.Username = "哈哈"
		user.Age = 40
		//保存所有字段
		models.DB.Save(&user)
		fmt.Println(user)
	*/
	user := models.User{}
	models.DB.Model(&user).Where("id=3").Update("username", "haha")

	c.String(http.StatusOK, "修改用户")

	//条件更新
}

func (con UserController) Delete(c *gin.Context) {
	user := models.User{Id: 3}
	models.DB.Delete(&user)
	c.String(http.StatusOK, "删除成功")
}

/*
1、获取上传的文件
2、获取后缀名判断类型是否正确.jpg .png .gif .jpeg
3、创建图片保存目录static/upload /20200623
4、生成文件名称和文件保存的目录
5、执行上传
*/
func (user UserController) DoUpload(c *gin.Context) {
	nickname := c.PostForm("nickname")
	// 1、获取上传的文件
	file, err := c.FormFile("file")

	dst := path.Join("server_02/static/upload", file.Filename)
	if err == nil {
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
			c.String(http.StatusOK, "上传的文件类型不合法")
			return
		}
		// 3、创建图片保存目录static/upload /20200623
		day := models.GetDay()
		dir := "server_02/static/upload/" + day
		//该方法可以创建多层目录
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusOK, "MkDir失败")
			return
		}
		// 4、生成文件名称和文件保存的目录
		unix := time.Now().Unix()
		filename := fmt.Sprintf("%v%v", unix, extName)

		//5、执行上传
		//路径拼接
		dst := path.Join(dir, filename)

		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"nickname": nickname,
		"dst":      dst,
	})

}
func (user UserController) DoUploadSame(c *gin.Context) {
	nickname := c.PostForm("nickname")
	//根据form提交的名字获取，类似PostForm
	form, err := c.MultipartForm()
	files := form.File["file[]"]
	dsts := make([]string, 0)
	for _, file := range files {
		//路径拼接
		dst := path.Join("server_02/static/upload", file.Filename)
		if err == nil {
			c.SaveUploadedFile(file, dst)
			dsts = append(dsts, dst)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"nickname": nickname,
		"dst":      dsts,
	})

}
func (user UserController) DoEdit(c *gin.Context) {
	nickname := c.PostForm("nickname")
	//根据form提交的名字获取，类似PostForm
	file1, err1 := c.FormFile("file1")
	//路径拼接
	dst1 := path.Join("server_02/static/upload", file1.Filename)
	if err1 == nil {
		c.SaveUploadedFile(file1, dst1)
	}
	file2, err2 := c.FormFile("file2")
	//路径拼接
	dst2 := path.Join("server_02/static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"nickname": nickname,
		"dst1":     dst1,
		"dst2":     dst2,
	})

}
