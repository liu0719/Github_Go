package dao

import (
	"fmt"
	"go_web/bookstore/model"
	"testing"
)

//	func TestBook(t *testing.T) {
//		fmt.Println("测试bookdao方法")
//		t.Run("", textBook)
//	}
func TestUser(t *testing.T) {
	fmt.Println("开始测试UserDao的函数")
	// t.Run("测试根据价格获取页数", testGetPageBoooksByPrice)
	// t.Run("测试获取页数", testGetPageBoooks)
	// t.Run("测试更新一本书", testUpdateBook)
	// t.Run("测试getBookById", testGetBookById)
	// t.Run("测试AddBook", testAddBook)
	// t.Run("测试DeleteBook", testDeleteBook)
	// t.Run("测试GetBook",testGetBook)
	// t.Run("验证用户名和密码", testLogin)
	// t.Run("验证用户名", testRegist)
	// t.Run("保存用户名和密码", testSave)
	// t.Run("添加session", testAddSession)
	// t.Run("删除session", testDeleteSession)
	// t.Run("获取session", testGetSession)
	// t.Run("添加到购物列表", testAddCart)

}

func testGetSession(t *testing.T) {
	sess, err := GetSession("52fdfc07-2182-454f-563f-5f0f9a621d72")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("session 的信息:", sess)
}
func testDeleteSession(t *testing.T) {

	err := DeleteSession("472342941")
	if err != nil {
		panic(err)
	}
}
func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionId: "472342941",
		UserName:  "haha",
		UserId:    8,
	}
	err := AddSession(sess)
	if err != nil {
		panic(err)
	}
}
func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123")
	fmt.Println("获取到的用户信息是", user)

}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取到的用户信息是", user)
}
func testSave(t *testing.T) {
	_ = SaveUser("admin99999", "321", "321@qq.com")
}
func testGetBook(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range books {
		fmt.Printf("第%v图书的信息是:%v\n", i+1, v)
	}
}
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   77.77,
		Sales:   101,
		Stock:   99,
		ImgPath: "static/img/default.jpg",
	}
	//调用添加函数
	err := AddBook(book)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func testDeleteBook(t *testing.T) {
	//调用删除函数
	err := DeleteBook(32)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func testGetBookById(t *testing.T) {
	book, err := GetBookById("30")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("获取图书信息是", book)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		Id:      31,
		Title:   "孙悟空",
		Author:  "吴",
		Price:   66.66,
		Sales:   1000,
		Stock:   9,
		ImgPath: "static/img/default.jpg",
	}
	err := UpdateBook(book)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetPageBoooks(t *testing.T) {
	page, err := GetPageBooks("5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("当前页属实：", page.PageNo)
	fmt.Println("当前zong：", page.TotalRecord)
	fmt.Println("当前页属实：", page.TotalPageNo)

	for _, v := range page.Books {
		fmt.Println(v)
	}
}
func testGetPageBoooksByPrice(t *testing.T) {
	page, err := GetPageBooksByPrice("1", "10", "30")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("当前页属实：", page.PageNo)
	fmt.Println("当前zong：", page.TotalRecord)
	fmt.Println("当前页属实：", page.TotalPageNo)

	for _, v := range page.Books {
		fmt.Println(v)
	}
}

func TestAddCart(t *testing.T) {
	fmt.Println("开始测试购物车相关函数")
	// t.Run("测试添加购物车", testGetCartItemByBookId)
	// t.Run("获取购物车中所有的购物项", testGetCartItemByCartId)
	// t.Run("根据UserId获取购物车", testGetCartByUserId)
	// t.Run("测试更新cartItem中的Count", testUpdateBookCount)
	// t.Run("测试删除购物车", testDeleteCartById)
	t.Run("测试删除购物车", testDeleteCartItemById)
	// t.Run("测试添加购物车", testAddCart)
}

func testAddCart(t *testing.T) {
	//设置要买的第一本书
	book1 := &model.Book{
		Id:    2,
		Price: 27.20,
	}
	//设置要买的第二本书
	book2 := &model.Book{
		Id:    3,
		Price: 23.00,
	}
	// 创建两个购物项
	cartitem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartId: "666888",
	}
	cartitem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartId: "666888",
	}
	//创建一个购物想切片
	var cartItemSlice []*model.CartItem
	cartItemSlice = append(cartItemSlice, cartitem1)
	cartItemSlice = append(cartItemSlice, cartitem2)
	//创建购物车
	cart := &model.Cart{
		CartId:    "666888",
		CartItems: cartItemSlice,
		UserId:    4,
	}
	err := AddCart(cart)
	if err != nil {
		panic(err)
	}
}

func testGetCartItemByBookId(t *testing.T) {
	cartItem, err := GetCartItemByBookIdAndCartId("2", "666888")
	if err != nil {
		panic(err)
	}
	fmt.Println("查询到的数据：", cartItem)
}

func testGetCartItemByCartId(t *testing.T) {
	cartSlice, err := GetCartItemByCartId("666888")
	if err != nil {
		panic(err)
	}
	for i, v := range cartSlice {
		fmt.Printf("第%v个购物项:%v\n", i, v)
	}
}

func testGetCartByUserId(t *testing.T) {
	cart, err := GetCartByUserId(4)
	if err != nil {
		panic(err)
	}
	fmt.Println("购物车信息:", cart)
}

func testUpdateBookCount(t *testing.T) {
	book := &model.Book{
		Id: 1,
	}
	cartItem := &model.CartItem{
		Count:  2,
		Book:   book,
		CartId: "2e225ec7-b23e-43f3-5758-1f3776083a18",
	}
	err := UpDateBookCount(cartItem)
	if err != nil {
		panic(err)
	}
}

func testDeleteCartById(t *testing.T) {
	err := DeleteCartById("a8f5090e-2732-4044-647f-c8afd0f9951c")
	if err != nil {
		panic(err)
	}
}
func testDeleteCartItemById(t *testing.T) {
	err := DeleteCartItemById("133")
	if err != nil {
		panic(err)
	}
}
func TestOrder(t *testing.T) {
	fmt.Println("测试订单的函数")
	// t.Run("测试添加订单和订单项", testAddOrder)
	// t.Run("测试获取所有订单项", testGetOrders)
	// t.Run("测试获取订单项", testGetOrderItemByOrderId)
	// t.Run("测试获取订单项", testGetOrdersByUserId)
	t.Run("测试更新状态", testUpdateOrderState)
}

func testUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("12345678", 1)
	if err != nil {
		panic(err)
	}
}
func testAddOrder(t *testing.T) {
	orderId := "12345678"
	// 创建一个订货单
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  "2003.01.01 12:12:12",
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserId:      1,
	}
	orderItem := &model.OredrItem{
		Count:   1,
		Amount:  100,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderId,
	}
	orderItem2 := &model.OredrItem{
		Count:   1,
		Amount:  300,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderId,
	}
	err := AddOrder(order)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = AddOrderItem(orderItem)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = AddOrderItem(orderItem2)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func testGetOrders(t *testing.T) {
	orders, err := GetOrders()
	if err != nil {
		panic(err)
	}
	for i, v := range orders {
		fmt.Println("第", i, "个信息是:", v)
	}
}
func testGetOrderItemByOrderId(t *testing.T) {
	orderItemSlice, err := GetOrderItemByOrderId("12345678")
	if err != nil {
		panic(err)
	}
	for _, v := range orderItemSlice {
		fmt.Println("订单像是:", v)
	}
}
func testGetOrdersByUserId(t *testing.T) {
	orderItemSlice, err := GetOrdersByUserId(9)
	if err != nil {
		panic(err)
	}
	for _, v := range orderItemSlice {
		fmt.Println("订单像是:", v)
	}
}
