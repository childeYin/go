go登录以及保持连接

因为我跟别人说，我的tcp/ip还没有看完，然后他说，你去写一个登录已经保持连接，基本能懂点，我去，就这么简单的给自己找了一个活，然后我说用啥语言，他说c,我就望着天空开始流泪~~后来他说，你挑吧，我就果断的选择了go,原因很简单，没玩过。

思路：

	1.go 接受用户输入的帐号密码，登录并保持连接？？？

	2.go 如何接受终端输入信息
		- bufio.NewReader

	3.go 操作数据库
		- database/sql
		- go get github.com/go-sql-driver/mysql

	4.go socket? or other? 如何保持连接

		- https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/08.1.md
		- https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
		- https://tour.golang.org/

	5. 两个client如何交互信息

	6. goroutine 没有进行完，main就结束了

	目前还有超多问题 但是我又近了一步

		1.多帐号登录 index out of range 错误；messages.go:15 [ok]
		2.掉线咋办 
		3.消息丢失 [ok]
		4.安装到服务器上
		5.server没有启动，client登录崩溃[ok]
		6.一端退出，另一端发送消息，服务器报错

		7.拆 注意模块化 [ok] ****
		8.在线功能做好了没？怎么检测某用户掉线了？其他用户怎么知道这用户掉线了，如果不知道发消息会怎样？[ok]
		9.上线通知[ok]
		10.解决重复登录问题[ok]

运行方式:

	1. ./make.sh
	2. ./login ./server

instruction :

	1)发送消息 msg;收信人;消息内容 (msg;可写,可不写)
	2)退出登录 quit
	3)查找好友 search;好友名字
