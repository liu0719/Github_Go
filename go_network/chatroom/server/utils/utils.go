package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	message "go_network/chatroom/common/messgae"
	"io"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct { //传输者
	//分析字段
	C   net.Conn
	Buf [8096]byte
}

//读取发送的信息
func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	_, err = t.C.Read(t.Buf[:4])
	if err != nil {
		if err == io.EOF {
			err = errors.New("客户端断开连接")
			return
		}
		fmt.Println("读取数据错误:", err)
		return
	}
	//根据buf[:4]转成yigeuint32类型
	pkgLen := binary.BigEndian.Uint32(t.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := t.C.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println(err)
		return
	}

	//将buf反序列化
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (t *Transfer) WriterPkg(data []byte) (err error) {

	//先发送长度给对方
	pkgLen := uint32(len(data))
	//将长度转成bytes
	// var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)
	n, err := t.C.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送失败:", err)
		return
	}

	//发送data,本身
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)
	n, err = t.C.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送失败:", err)
		return
	}
	return
}
