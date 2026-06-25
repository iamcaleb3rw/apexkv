package main

import (
  "fmt"
  "log"
  "net"
  "os"
)


func main(){
	log.Println("reading config file")
	readConf("./redis.conf")


	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal("Cannot Listen on :6379")
	}

	defer l.Close()
    log.Println("Listening on :6379")
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Listening on :6379")

	defer conn.Close()

	log.Println("Connection eccepted")
    
	for {
        v:= Value{typ: ARRAY}
		v.readArray(conn)
		handle(conn, &v)

		fmt.Println(v.array)
	}
	
}




