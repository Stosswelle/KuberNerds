import threading

import requests
import time 

IP = '127.0.0.1'
PORT = '38955'
PAGE = 'productpage'
URL = "http://{}:{}/{}".format(IP, PORT, PAGE)

TIMES = 5
THREADS = 10

def millis():
    return int(round(time.time() * 1000))

def get():
    start_time = millis()
    result = requests.get(URL)
    end_time = millis()
    print(end_time - start_time)
    result.close()
    return 0

class Client(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)

    def run(self):
        for i in range(TIMES):
            get()

def main():
    for i in range(THREADS):
        new_thread = Client()
        new_thread.start()

main()
