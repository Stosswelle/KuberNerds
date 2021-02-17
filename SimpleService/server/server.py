from flask import Flask, request
app = Flask(__name__)

TCP_PORT = 12235

@app.route('/')
def hello():
    hello = request.form.get('hello')
    if hello: 
        return "Hello Leiyi!"
    else:
        return 'Hellon\'t Leiyi'

if __name__ == '__main__':
    app.run(debug=True, port=TCP_PORT)
