import socket

SERVER_ADDR = '127.0.0.1'
TCP_PORT = 12235

def main():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect((SERVER_ADDR, TCP_PORT))

    message = b'Hello Andrew!\0'
    client.send(message)

    reply = client.recv(10)
    print("Received message: " + reply.decode())

    client.close()

    return 0

main()