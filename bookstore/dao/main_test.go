package dao

import (
	"bookstore/models"
	"fmt"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	// t.Run("测试登录", testLogin)
	// t.Run("测试检测用户名是否已存在", testRegist)
	// t.Run("测试保存到数据库", testSaveUser)
}
func testLogin(t *testing.T) {
	user, err := CheckUserNameAndPassword("admin", "123")
	if err != nil {
		fmt.Println("eee", err)
		return
	}
	fmt.Println(user)
}
func testRegist(t *testing.T) {
	user, err := CheckUserName("admin10")
	if err != nil {
		fmt.Println("eee", err)
		return
	}
	fmt.Println(user)
}
func testSaveUser(t *testing.T) {
	err := SaveUser("admin1", "123", "123@126.com")
	if err != nil {
		fmt.Println(err)
	}
}
func TestBooks(t *testing.T) {
	// t.Run("根据分页获取图书", testGetPageBooks)
	// t.Run("删除图书", testDeleteBook)
	t.Run("添加图书", testAddBook)
	// t.Run("更新图书", testUpDateBook)
}

func testDeleteBook(t *testing.T) {
	err := DeleteBook(33)
	if err != nil {
		fmt.Println(err)
	}
}
func testAddBook(t *testing.T) {
	book := models.Book{
		Id:      33,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   88.88,
		Sales:   100,
		Stock:   999,
		ImgPath: "/static/img/三体.jpg",
	}
	err := AddBook(&book)
	if err != nil {
		fmt.Println(err)
	}
}
func testUpDateBook(t *testing.T) {
	book := models.Book{
		Id:      33,
		Title:   "西游记www",
		Author:  "吴承恩",
		Price:   88.88,
		Sales:   100,
		Stock:   999,
		ImgPath: "/static/img/三体.jpg",
	}
	err := AddBook(&book)
	if err != nil {
		fmt.Println(err)
	}
}
func testGetPageBooks(t *testing.T) {
	page, err := GetPageBooks("5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(page)
	for _, v := range page.Books {
		fmt.Println(v)
	}
}

func TestGetSession(t *testing.T) {
	// t.Run("根据SessionID获取session", testGetSession)
	// t.Run("添加session", testAddSession)
	t.Run("删除session", testDeleteSession)
}
func testGetSession(t *testing.T) {
	session, err := GetSession("123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)
}
func testAddSession(t *testing.T) {
	session := &models.Session{
		UserId:    2,
		Username:  "admin2",
		SessionId: "321",
	}
	err := AddSession(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("添加成功--->", session)
}
func testDeleteSession(t *testing.T) {
	err := DeleteSession("321")
	if err != nil {
		fmt.Println("eeee", err)
	}
	fmt.Println("删除成功")
}
func TestCart(t *testing.T) {
	// t.Run("GetCartItemByCartId", testGetCartItemByCartId)
	// t.Run("testGetCartByUserId", testGetCartByUserId)
	// t.Run("testGetCartItemByBookIdAndCartId", testGetCartItemByBookIdAndCartId)
	// t.Run("testUpDateBookCount", testUpDateBookCount)
	// t.Run("UpDateCart", testUpDateCart)
	// t.Run("AddCart", testAddCart)

	t.Run("testAddCartItem", testAddCartItem)

}
func testAddCart(t *testing.T) {
	var cartItems []*models.CartItem
	book1 := &models.Book{
		Id:      35,
		Title:   "三国演义",
		Author:  "刘广硕",
		Price:   12.34,
		Sales:   999,
		Stock:   101,
		ImgPath: "/static/img/default.jpg",
	}
	cartItem1 := models.CartItem{
		Book:   book1,
		BookId: 33,
		Count:  1,
		CartId: "123",
	}
	cartItem1.Amount = cartItem1.GetAmount()
	book2 := &models.Book{
		Id:      36,
		Title:   "西游记",
		Author:  "刘广硕",
		Price:   12.34,
		Sales:   101,
		Stock:   999,
		ImgPath: "/static/img/default.jpg",
	}
	cartItem2 := models.CartItem{
		Book:   book2,
		BookId: 36,
		Count:  1,
		CartId: "123",
	}
	cartItem1.Amount = cartItem1.GetAmount()
	cartItems = append(cartItems, &cartItem1)
	cartItems = append(cartItems, &cartItem2)
	cart := &models.Cart{
		Id:         "123",
		CartItems:  cartItems,
		TotalCount: 2,
		UserId:     1,
	}
	cart.TotalAmount = cart.GetTotlalAmount()
	err := AddCart(cart)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("成功")
}
func testGetCartItemByCartId(t *testing.T) {
	cartItems, err := GetCartItemByCartId("123")
	if err != nil {
		fmt.Println(err)
	}
	if cartItems == nil {
		fmt.Println(cartItems)
	}

	for _, v := range cartItems {
		fmt.Println(v)
		fmt.Println(v.Book)
	}
}

func testGetCartByUserId(t *testing.T) {
	cart, err := GetCartByUserId(9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cart)
	for _, v := range cart.CartItems {
		fmt.Println(v)
		fmt.Println("书", v.Book)
	}
}

func testGetCartItemByBookIdAndCartId(t *testing.T) {
	cartItem, err := GetCartItemByBookIdAndCartId("1", "0")
	if err != nil {
		fmt.Println("err====", err)
	}
	fmt.Println(cartItem)
	fmt.Println(cartItem.Book)

}

func testUpDateBookCount(t *testing.T) {
	cartItem := &models.CartItem{
		Id:     1,
		Count:  999,
		Amount: 999,
		BookId: 1,
		CartId: "123",
	}
	err := UpDateBookCount(cartItem)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cartItem)

}

// func testUpDateCart(t *testing.T) {
// 	cart := &models.Cart{
// 		Id:          "123",
// 		TotalCount:  99,
// 		TotalAmount: 99,
// 		UserId:      1,
// 	}

//		err := UpDateCart(cart)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func testAddCartItem(t *testing.T) {
	book1 := &models.Book{
		Id:      35,
		Title:   "三国演义",
		Author:  "刘广硕",
		Price:   12.34,
		Sales:   999,
		Stock:   101,
		ImgPath: "/static/img/default.jpg",
	}
	cartItem := &models.CartItem{
		CartId: "123",
		Count:  7,
		Amount: 7,
		BookId: 4,
		Book:   book1,
	}

	err := AddCartItem(cartItem)
	if err != nil {
		fmt.Println(err)
	}
}
func TestOrder(t *testing.T) {
	t.Run("添加Order", testAddOrder)
	// t.Run("获取Order", testGetOrders)
}
func testAddOrder(t *testing.T) {
	order := &models.Order{
		Id:          "1234sfs53t34t",
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  1,
		TotalAmount: 1,
		State:       1,
		UserId:      1,
	}
	err := AddOrder(order)
	if err != nil {
		fmt.Println(err)
	}
}
func TestGetOrders(t *testing.T) {
	orders, err := GetOrders()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orders)
	for _, v := range orders {
		fmt.Println(v)
	}
}
func TestGetOrdersByUserId(t *testing.T) {
	orders, err := GetOrdersByUserId(2)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(orders)
	for _, v := range orders {
		fmt.Println(v)
	}
}
func TestUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("1234sfs53t34t", 2)
	if err != nil {
		fmt.Println(err)
	}
}
