# hysteria2-pinSHA256

## Sev00、Ct8直接部署
```bash
bash <(curl -s https://128877.xyz/sha256.sh)
```
调用
```
https://webapi.serv00.net/sha256.php?1.1.1.1:0
```
返回结果
```bash
#不带查询字符串时
ERROR!请输入IP和端口号<   ?IP:PORT   >
#主机不通时
ERROR
#正常返回：固定证书的sha256
b4f7bc5565fea3521dbf040e143b065c6f9bbd85d62d8f90e7494d2eadda899c
```
节点修复
```
#无法获取固定证书时，内核使用sing-box
hysteria2://UUID@IP:PORT?sni=www.bing.com&insecure=1#TITLE
#获取固定证书后:CERT
hysteria2://UUID@IP:PORT?sni=www.bing.com&insecure=0&allowInsecure=0&pinSHA256=CERT#TITLE
```
