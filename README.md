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

	目前还有超多问题 但是我又近了一步