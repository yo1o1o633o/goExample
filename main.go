package main

// 不会有输出, 因为主进程已经退出了
func main() {
	go func() {
		println("Hello")
	}()

	go func() {
		println("GO")
	}()
}