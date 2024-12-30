package main

func main() {

	server := NewAPIServer(":7676")
	server.Run()

	//fmt.Println("Quick Test")
}
