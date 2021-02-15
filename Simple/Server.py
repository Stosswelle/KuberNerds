import socket
import struct
from threading import Thread

TCP_PORT = 12235

class Server(Thread):
    def __init__(self, client_socket):
        Thread.__init__(self)
        self.client_socket = client_socket

    def run(self):
        message = self.client_socket.recv(14)
        print("Received message: " + message.decode())
            
        reply = b'Hi Leiyi!\0'
        self.client_socket.send(reply)
        return 0


def main():
    print("Listenning on port " + str(TCP_PORT) + "...")
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind(('', TCP_PORT))
    threads = []

    while 1:
        server.listen()
        try:
            client_socket, client_addr = server.accept()
        except:
            print("Connection Error!")
            break

        print("Connection setup!")
        print("Client Address: " + str(client_addr))
        new_server = Server(client_socket)
        new_server.start()
        threads.append(new_server)
    
    server.close()
    for t in threads:
        t.join()

main()
