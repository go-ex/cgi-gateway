## 启动参数

./gateway addr=":8080" // "服务端口配置

./gateway conf="./config.json" // 配置文件路径

## websocket

host:8080/  websocket提供服务端口

发起连接,默认按 app_id=1 转发到后台连接

    ws://127.0.0.1:8080?app_id=1 

# config

    {
      "apps": [
        {
          "id":1,
          "url":"http://local-lrs-tt.7955.com/ws"
        }
      ]
    }


## api
host:8080/app               应用操作
host:8080/app/add           应用操作
host:8080/config            配置显示
host:8080/broadcast         广播所有
    {"msg":""}
host:8080/ws/close   关闭指定连接
    {"fd":""}
host:8080/ws/send    发送信息
    {"fd":"","msg":""}


