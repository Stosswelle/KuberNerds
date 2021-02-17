import requests

# Address and port number for load balancer
SERVER_ADDR = '127.0.0.1'
TCP_PORT = 46715 # 12235 for local testing, port depends on what is exposed for lb
TIMES = 1

def main():
    address = 'http://{SERVER_ADDR}:{TCP_PORT}'.format(SERVER_ADDR=SERVER_ADDR, TCP_PORT=TCP_PORT)
    for i in range(TIMES):
        payload = {'hello': 'Jeff'}
        result = requests.get(address, params=payload)
        print(result.status_code, result.content)
        result.close()
    return 0

main()
