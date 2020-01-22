## Hello Internet in Go

### Team Memebers

- Guodong Xie
- HeNian Wang
- Samuel Farid


### How To Build

'''
go build client.go
go build server.go
'''

### How to Run

'''

# for server, you need to provide HOST

./server PORT

# for client, you need to provide both HOST and PORT

./clinet HOST PORT

'''

### Key Concepts

#### Listen and Accept

Use the `net.Listen("tcp", ":"+PORT)` function to start a service on PORT and use `instance.Accept()` to accpet new connections.

#### Establishing a connection

Use the `net.Dial(HOST, PORT)` function to create a TCP socket in python. This will return a Connnection object.

#### Sending

Use the `fmt.Fprintf(conn, Message)` to send messages in a connection

#### Sending and Receiving

Use the `bufio.NewReader(conn).ReadString('\n')` to read messages

#### Clean up

You should `conn.Close()` a socket when you are done with it.
