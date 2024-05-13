package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type Transfer struct {
	C   net.Conn
	Buf [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:10], pkgLen)
	n, err := this.C.Write(this.Buf[:10])
	if err != nil || n != 10 {
		fmt.Println("发送失败:", err)
		return
	}

	n, err = this.C.Write(data)
	if err != nil {
		fmt.Println("发送失败:", err)
	}

	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	n, err := this.C.Read(this.Buf[:10])
	if err != nil || n != 10 {
		if err == io.EOF {
			err = errors.New("连接已断开")
			return
		}
		fmt.Println("服务器接收异常参数:", err)
		return
	}
	pkgLen := binary.BigEndian.Uint32(this.Buf[:10])
	n, err = this.C.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("饭序列化失败:", err)
	}
	return
}
