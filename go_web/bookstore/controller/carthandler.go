package controller

import (
	"encoding/json"
	"fmt"
	"go_web/bookstore/dao"
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
)

type CartHandler struct {
	CartBool bool
	Cart     *model.Cart
	Session  *model.Session
}

// 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemId, err := strconv.ParseInt(cartItemId, 10, 64)
	if err != nil {
		return
	}
	_, session := dao.IsLogin(r)
	userId := session.UserId
	// 获取到该用户Id对应的购物车
	cart, _ := dao.GetCartByUserId(userId)

	cartItemSlice := cart.CartItems
	for i, v := range cartItemSlice {
		if v.CartItemId == iCartItemId {
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
		err := dao.DeleteCartById(cart.CartId)
		if err != nil {
			return
		}
		GetCartInfo(w, r)
		return
	}
	if err != nil {
		return
	}
	//调用GetCartInfo再次查询购物车
	GetCartInfo(w, r)
}

// 删除购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	err := dao.DeleteCartById(cartId)
	if err != nil {
		return
	}
	//调用GetCartInfo再次查询购物车
	GetCartInfo(w, r)
}

// 根据用户ID获取购物信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	//获取用户的ID
	userId := session.UserId
	//根据用户ID获取该ID对应的购物车
	cart, _ := dao.GetCartByUserId(userId)

	if cart != nil {

		cart.UserName = session.UserName
		//解析模板
		t := template.Must(template.ParseFiles("bookstore/views/pages/cart/cart.html"))
		cartHandler := &CartHandler{
			CartBool: true,
			Cart:     cart,
		}
		t.Execute(w, cartHandler)

	} else {
		//该用户还没有购物车
		cart := &model.Cart{
			UserName: session.UserName,
			UserId:   session.UserId,
		}
		cartHandler := &CartHandler{
			CartBool: false,
			Session:  session,
			Cart:     cart,
		}
		t := template.Must(template.ParseFiles("bookstore/views/pages/cart/cart.html"))
		t.Execute(w, cartHandler)
	}

}

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//获取前端传来的图书ID
	bookId := r.FormValue("bookId")
	// 根据图书ID获取图书信息
	book, err := dao.GetBookById(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("获取图书的信息时", book)
	// 判断是否登录
	flag, session := dao.IsLogin(r)
	if !flag {
		w.Write([]byte("请先登录！"))
		return
	}
	// 获取用的Id
	userId := session.UserId
	//判断当前数据库中是否有当前用户选择的书籍
	cart, _ := dao.GetCartByUserId(userId)
	//设置返回时的的用户名

	if cart != nil {
		//说明该用户已经有购物车，此时需要判断购物车中是否由当前这本图书
		cartItem, _ := dao.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
		// fmt.Println("根据图书Id获取的购物项是:", cartItem)
		if cartItem != nil {
			//购物车中已经由该图书，将count+1就可以
			// 获取购物车切片所有中的购物项
			cartItemSlice := cart.CartItems
			for _, v := range cartItemSlice {
				//找到和要添加的图书相对应的购物项
				if v.Book.Id == cartItem.Book.Id {
					//相等表示为同一本书，
					v.Count++
					dao.UpDateBookCount(v)
					w.Write([]byte(fmt.Sprintf("您刚刚将《%v》* %v加入到了购物车", v.Book.Title, v.Count)))
				}

			}
		} else {
			// 购物车中没有这个图书，此时创建购物项瓶添加到数据库
			//创建购物车中的购物项
			newcartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartId: cart.CartId,
			}
			//将购物项放入到购物想切片中
			//将切片放入到当前已经有的购物车中
			cart.CartItems = append(cart.CartItems, newcartItem)
			//将新创建的购物项添加到数据库中
			err := dao.AddCartItem(newcartItem)
			if err != nil {
				return
			}
			str := "您刚刚将《" + book.Title + "》加入到了购物车"
			w.Write([]byte(str))
		}

		//不管之前购物车中是否有图书对应得购物项，购物车的count都必须更新总数量和总金额
		err := dao.UpDateCart(cart)
		if err != nil {
			return
		}
	} else {
		//说明该用户还没有购物车
		//没有购物车加就需要为该用户创建一个购物车
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartId: cartId,
			UserId: userId,
		}
		//创建购物车中的购物项
		var cartItemSlice []*model.CartItem

		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartId: cartId,
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
		str := "您刚刚将《" + book.Title + "》加入到了购物车"
		w.Write([]byte(str))
	}

}

// 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	// 获取用户输入的值
	cartItemId := r.FormValue("cartItemId")
	bookCount := r.FormValue("bookCount")
	ibookCount, err := strconv.ParseInt(bookCount, 10, 64)
	if err != nil {
		return
	}

	iCartItemId, err := strconv.ParseInt(cartItemId, 10, 64)
	if err != nil {
		return
	}
	_, session := dao.IsLogin(r)
	userId := session.UserId
	// 获取到该用户Id对应的购物车
	cart, _ := dao.GetCartByUserId(userId)

	cartItemSlice := cart.CartItems
	for _, v := range cartItemSlice {
		if v.CartItemId == iCartItemId {
			//这个就是我们要更新得购物项
			//将当前的数量设置为用户输入的数量
			v.Count = ibookCount
			//更新数据库中该购物项的金额和数量
			dao.UpDateBookCount(v)
			break
		}
	}
	if ibookCount == 0 {
		err := dao.DeleteCartById(cart.CartId)
		if err != nil {
			return
		}
		GetCartInfo(w, r)
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
		if iCartItemId == v.CartItemId {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
