package phchat
import (
    "bufio"
    "fmt"
    "log"
    "net"
)
 type client chan<- string // Define a channel that sends data outward
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
    clients := make(map[client]bool) // Store the login status of each client
    for {
        select {
        case msg := <-messages:
            for cli := range clients {
                cli <- msg
            }
        case cli := <-entering:
            clients[cli] = true
        case cli := <-leaving:
            delete(clients, cli)
            close(cli)
        }
    }
}
 func handleConn(conn net.Conn) {
    ch := make(chan string)
    go clientWriter(conn, ch)
    who := conn.RemoteAddr().String()
    ch <- "You are " + who + "\n"
    entering <- ch
    messages <- who + " has arrived"
    fmt.Println(who + " has connected") // Print the user's connection to the console
    input := bufio.NewScanner(conn)
    for input.Scan() {
        messages <- who + ":" + input.Text() + "\n"
        fmt.Println(who + ":" + input.Text()) // Print the message to the console
    }
    leaving <- ch
    messages <- who + " has left"
    fmt.Println(who + " has disconnected") // Print the user's disconnection to the console
    conn.Close()
}
 func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg)
    }
}