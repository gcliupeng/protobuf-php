package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"net"
	"bytes"
    "encoding/binary"
	"github.com/golang/protobuf/proto"
	"pb"
	"bufio"
)

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
    
    rb:=bufio.NewReader(conn)
    for{
    	buf := make([]byte, 1000)
    	//四个字节存储data长度
		for i:=0;i<4;i++ {
			b,err:=rb.ReadByte()
		if(err!=nil){
				fmt.Printf("%v",err)
				return
			}
			buf[i]=b
		}
    	var x int32  
    	b_buf  :=  bytes .NewBuffer(buf)
    	err:=binary.Read(b_buf, binary.BigEndian, &x) 
    	fmt.Printf("get msg leng :%d\n",x)   
    	if err != nil {
        	// if err == io.EOF {                                                                                         
        	// 	return 
        	//           }
        	return 
    	}
    	len:=0
    	//read byte by byte
    	for  {
    		b,_:=rb.ReadByte()
    		buf[len]=b
    		len=len+1
    		if(int32(len)==x){
    			break
    			}
        	}
    	book := &pb.AddressBook{}
		if err := proto.Unmarshal(buf[:x], book); err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}
		listPeople(os.Stdout, book)
	}
}

// Main reads the entire address book from a client and prints all the
// information inside.
func main() {

	ln, err := net.Listen("tcp", ":9872")   
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

