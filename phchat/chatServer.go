package phchat
import (
    "bufio"
    "log"
    "net"
)
type Tuple struct {
  name string 
	cn   chan string
}
 type client Tuple
var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string)
)
 func StartServer() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal("network is broken", err)
    }
    log.Print("connect success!")
    go broadcaster()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}
 func broadcaster() {
	clients := make(map[client]bool)
    for {
        select {
        case msg := <-messages:
            for cli := range clients {
                cli.cn <- msg
            }
			log.Println(msg)
        case cli := <-entering:
			log.Println("connected "+cli.name)
            clients[cli] = true
        case cli := <-leaving:
			log.Println("disconnected" + cli.name)
            delete(clients, cli)
			close(cli.cn)
        }
    }
}
 func handleConn(conn net.Conn) {
    cli := client{}
	cli.cn = make(chan string)
    go clientWriter(conn, cli.cn)
    who := conn.RemoteAddr().String()
    cli.name = "You are " + who + "\n"
    cli.cn <- "You are " + who + "\n"
	entering <- cli
    messages <- who + " has arrived"
    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ":" + input.Text() + "\n"
    }
	leaving <- cli
    messages <- who + " has left"
    conn.Close()
}
 func clientWriter(conn net.Conn, ch chan string) {
    for msg := range ch {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}
    }
}