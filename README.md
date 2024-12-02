Sockets Programming Overview
A socket is a software abstraction that enables communication between two or more computer processes over a network. In this exercise, we will create two basic Go programs: a 'server' and a 'client'. These programs will communicate using a predefined protocol over a network. The client will send messages to the server, which will simply echo the received message back to the client.

Protocol: What is it?
Socket connections are established using a communication protocol (e.g., TCP, UDP) and the network address of the process with which they are connecting. A protocol is essentially a set of rules governing communication. In everyday life, we implicitly agree on a 'protocol' by using the same language to communicate.

For instance, imagine I open a car dealership (acting as the server) and a customer (the client) attempts to negotiate in German, a language I don’t speak. This is akin to a client attempting to communicate with a server using an unsupported protocol. Protocols can define restrictions such as message format, length, or content. This ensures clear, efficient communication between the client and server.

TCP (Transmission Control Protocol) and IP (Internet Protocol)
TCP is a widely used protocol for establishing socket connections, while IP governs how processes communicate over the internet. When using TCP/IP, the network address of a process is uniquely defined by an IP address and a port number. We will focus on TCP and IPv4 in this example (though IPv6 is available, IPv4 remains more commonly used).

For example, to set up a client program to communicate with a server at IP address 192.168.62.88 and port 5555, the full network address is 192.168.62.88:5555. If either the IP address or the port number is incorrect, communication will likely fail unless there is another process listening at the wrong combination.

An analogy: Think of a computer as a building. The IP address is the building’s address, and the port numbers are the flat numbers. To deliver a package, it's essential to know both the building and the flat number. If either is wrong, the package will be misdirected.

Client-Server Communication
To set up a socket connection, we must first choose a protocol (e.g., TCP) and understand that an IP address and port number define a network address. The next step is enabling communication between the client and server.

We can think of the client and server as two people on opposite sides of a seesaw. When one side goes up, the other must go down, and vice versa. The client sends a message, the server receives it, the server sends a response, and the client receives it.

Here’s how this works in practice:

Server: First, we create a server that listens for incoming connections. It needs to open a socket, specify a protocol (TCP), and use a port number (e.g., 5555).
Client: The client creates its socket using the same protocol (TCP) and provides the server’s network address (IP and port).
Connection: The client "dials" the server to establish a connection. Once the server "accepts" the connection, communication can begin.
Messaging: Once the connection is established, the client and server take turns sending and receiving messages. In this case, the client will first send a message (e.g., the user’s name) to the server, and the server will echo it back to the client.
For this exercise, the server will echo whatever message the client sends, and the client will display the echoed message. User input for the message will be gathered from the terminal using os.Stdin, and messages will be stored as a slice of bytes since that's how computers store data.

Summary of Steps:
Server: Create a socket connection, specifying a protocol (TCP) and port number.
Client: Create a socket connection, specifying a protocol (TCP) and the server’s network address (IP:port).
Server: Listen for incoming client connections.
Client: Attempt to connect to the server.
Server: Accept the client connection.
Loop: The client writes a message and the server reads it. Then, the server writes a message and the client reads it.
This back-and-forth communication forms the basis of socket programming.
