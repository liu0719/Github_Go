package controllers

import (
	"bookstore/dao"
	"bookstore/models"
	"bookstore/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
}

var (
	uuid string
)

// 登录
func (ApiHandler) Login(c *gin.Context) {
	//判断是否登录
	flag, _ := dao.IsLogin(c)
	if flag {
		//已经等登录就调用
		c.Redirect(http.StatusFound, "/")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := dao.CheckUserNameAndPassword(username, password)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", "用户名或密码不正确!")
	} else {
		//登陆成功
		// 创建session
		// 创建唯一的SessionId
		uuid = utils.CreateUUID()
		session := &models.Session{
			SessionId: uuid,
			Username:  user.Username,
			UserId:    user.Id,
		}
		//将session保存到数据库
		dao.AddSession(session)
		//创建cookie与session相关联
		c.SetCookie("user", uuid, 3600, "/", ".cpolar.top", false, true)
		c.HTML(http.StatusOK, "login_success.html", user)
	}
}

// 注册
func (ApiHandler) Regist(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	user, err := dao.CheckUserName(username)
	//等于控说明查到数据了,有该用户
	if err == nil {
		c.HTML(http.StatusOK, "regist.html", "用户名已存在！")
	} else {
		dao.SaveUser(username, password, email)
		c.HTML(http.StatusOK, "regist_success.html", user)
	}

}

// 根据ID检测用户是否已经存在
func (ApiHandler) CheckUserName(c *gin.Context) {
	username := c.Request.PostFormValue("username")

	user, err := dao.CheckUserName(username)
	//等于控说明查到数据了,有该用户
	if err == nil && user != nil {
		c.String(http.StatusOK, "%v", "用户名已存在！")
	} else {
		c.String(http.StatusOK, "%v", "<font style=\"color:green;\">用户名可用</font>")
	}
}

// 登出
func (ApiHandler) Logout(c *gin.Context) {
	//获取cookie
	user, err := c.Cookie("user")
	if err != nil {
		return
	}
	c.SetCookie("user", uuid, -1, "/", ".cpolar.top", false, true)
	//删除数据库中与之对应的session
	err = dao.DeleteSession(user)
	if err != nil {
		return
	}
	c.Redirect(http.StatusFound, "/")
}

// 添加书籍到购物车
func (ApiHandler) AddBookCart(c *gin.Context) {
	bookId := c.PostForm("bookId")
	iBookId, _ := strconv.Atoi(bookId)
	//获取完整的图书信息
	book, err := dao.GetBookById(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	//判断是否登录，获取cookie和session里的信息
	flag, session := dao.IsLogin(c)
	if !flag {
		c.String(http.StatusOK, "请先登录！")
		return
	}
	//获取session里的userId
	userId := session.UserId
	cart, err := dao.GetCartByUserId(userId)
	if err == nil {
		//说明该用户存在购物车
		cartItem, err := dao.GetCartItemByBookIdAndCartId(bookId, cart.Id)
		if err == nil {
			//购物车中已经由该图书，将count+1
			for _, v := range cart.CartItems {
				//找到和要添加的图书相对应的购物项
				if v.Book.Id == cartItem.Book.Id {
					if v.Book.Id == cartItem.Book.Id {
						v.Count++
						//相等表示为同一本书，
						dao.UpDateBookCount(v)
					}
					c.String(http.StatusOK, "您刚刚将《%v》* %v加入到了购物车", v.Book.Title, v.Count)
				}
			}
		} else {
			// 购物车中没有这个图书，此时创建购物项瓶添加到数据库
			//创建购物车中的购物项
			newcartItem := &models.CartItem{
				Count:  1,
				CartId: cart.Id,
				BookId: iBookId,
				Book:   book,
			}

			//将购物项放入到购物想切片中
			//将切片放入到当前已经有的购物车中
			cart.CartItems = append(cart.CartItems, newcartItem)
			//将新创建的购物项添加到数据库中
			err := dao.AddCartItem(newcartItem)
			if err != nil {
				return
			}
			c.String(http.StatusOK, "您刚刚将《"+book.Title+"》加入到了购物车")
		}
		//不管之前购物车中是否有图书对应得购物项，购物车的count都必须更新总数量和总金额
		err = dao.UpDateCart(cart)
		if err != nil {
			return
		}
	} else {
		//说明该用户还没有购物车
		//没有购物车加就需要为该用户创建一个购物车
		cartId := utils.CreateUUID()
		cart := &models.Cart{
			Id:     cartId,
			UserId: userId,
		}
		//创建购物车中的购物项
		var cartItemSlice []*models.CartItem

		cartItem := &models.CartItem{
			Book:   book,
			Count:  1,
			CartId: cartId,
			BookId: iBookId,
		}

		//将购物项放入到购物想切片中
		cartItemSlice = append(cartItemSlice, cartItem)
		//将切片放入到购物车中
		cart.CartItems = cartItemSlice
		//将购物车添加到数据库中
		err := dao.AddCart(cart)
		if err != nil {
			return
		}
		c.String(http.StatusOK, "您刚刚将《"+book.Title+"》加入到了购物车")

	}
}

// 根据用户ID获取购物信息
func (ApiHandler) GetCartInfo(c *gin.Context) {
	_, session := dao.IsLogin(c)
	//获取用户的ID
	userId := session.UserId
	//根据用户ID获取该ID对应的购物车
	cart, err := dao.GetCartByUserId(userId)

	if err == nil {

		cart.User.Username = session.Username
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"CartBool": true,
			"Cart":     cart,
		})
	} else {
		//该用户还没有购物车
		cart := &models.Cart{
			User: &models.User{
				Username: session.Username,
				Id:       session.UserId,
			},

			UserId: session.UserId,
		}
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"CartBool": false,
			"Session":  session,
			"Cart":     cart,
		})
	}

}

// 删除购物车
func (ApiHandler) DeleteCart(c *gin.Context) {
	cartId := c.Query("cartId")
	err := dao.DeleteCartById(cartId)
	if err != nil {
		return
	}
	//调用GetCartInfo再次查询购物车
	ApiHandler{}.GetCartInfo(c)
}

// 删除购物项
func (ApiHandler) DeleteCartItem(c *gin.Context) {
	cartItemId := c.Query("cartItemId")
	iCartItemId, err := strconv.Atoi(cartItemId)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, session := dao.IsLogin(c)
	if session == nil {
		c.Redirect(http.StatusFound, "/api/login")
	}
	userId := session.UserId
	// 获取到该用户Id对应的购物车
	cart, _ := dao.GetCartByUserId(userId)

	cartItemSlice := cart.CartItems
	for i, v := range cartItemSlice {
		if v.Id == iCartItemId {
			//这个就是我们要删除得购物项
			// 将当前这个从切片中移除
			cartItemSlice = append(cartItemSlice[:i], cartItemSlice[i+1:]...)
			cart.CartItems = cartItemSlice
			// 将当前购物项从数据库中删除
			dao.DeleteCartItemById(cartItemId)
			dao.UpDateCart(cart)
			break
		}
	}
	if len(cart.CartItems) == 0 {
		err := dao.DeleteCartById(cart.Id)
		if err != nil {
			return
		}
		ApiHandler{}.GetCartInfo(c)
		return
	}
	if err != nil {
		return
	}
	//调用GetCartInfo再次查询购物车
	ApiHandler{}.GetCartInfo(c)
}

// 更新购物项
func (ApiHandler) UpdateCartItem(c *gin.Context) {
	// 获取用户输入的值
	cartItemId := c.PostForm("cartItemId")
	bookCount := c.PostForm("bookCount")
	ibookCount, err := strconv.Atoi(bookCount)
	if err != nil {
		return
	}

	iCartItemId, err := strconv.Atoi(cartItemId)
	if err != nil {
		return
	}
	_, session := dao.IsLogin(c)
	userId := session.UserId
	// 获取到该用户Id对应的购物车
	cart, _ := dao.GetCartByUserId(userId)

	cartItemSlice := cart.CartItems
	for _, v := range cartItemSlice {
		if v.Id == iCartItemId {
			//这个就是我们要更新得购物项
			//将当前的数量设置为用户输入的数量
			v.Count = ibookCount
			//更新数据库中该购物项的金额和数量
			dao.UpDateBookCount(v)
			break
		}
	}
	if ibookCount == 0 {
		err := dao.DeleteCartById(cart.Id)
		if err != nil {
			return
		}
		ApiHandler{}.GetCartInfo(c)
		return
	}
	dao.UpDateCart(cart)
	//查询购物车信息
	cart, _ = dao.GetCartByUserId(userId)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemId == v.Id {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"Amount":      amount,
		"TotalAmount": totalAmount,
		"TotalCount":  totalCount,
	})
}
