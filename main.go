package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const heart = `

     ***********                  ***********
  *****************            *****************
*********************        *********************
***********************      ***********************
************************    ************************
*************************  *************************
 **************************************************
  ************************************************
    ********************************************
      ****************************************
         **********************************
           ******************************
              ************************
                ********************
                   **************
                     **********
                       ******
                         **

`

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Esimiñ kim, sulu?: "))
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	lines := strings.Split(heart, "\n")
	for _, line := range lines {
		conn.Write([]byte(line + "\n"))
	}
	conn.Write([]byte(fmt.Sprintf("\nHappy Valentine's Day, %v! ❤️\n", name)))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP Server started on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}
