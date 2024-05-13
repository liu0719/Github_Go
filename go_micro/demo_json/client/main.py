import socket
import json

AddGoods={
    "id":10,
    "Title":  "西游记",
	"Author":  "没法说",
	"Context": "haahaha"
}
request = {
    "id": 0,
    "params": [AddGoods],
    "method": "goods.AddGoods"
}

client = socket.create_connection(("10.50.142.216", 80),5)
client.sendall(json.dumps(request).encode())

rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)

