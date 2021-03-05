from flask import Flask, request
app = Flask(__name__)

TCP_PORT = 12235

@app.route('/')
def hello():
    hello = request.args.get('hello')
    print(hello, request.args)
    if hello: 
        return "Hello Leiyi!"
    else:
        return 'Hellon\'t Leiyi'

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True, port=TCP_PORT)
