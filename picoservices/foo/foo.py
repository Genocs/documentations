from flask import request
from flask import Flask

#from flask import render_template

app = Flask(__name__)


@app.route('/')
def hello():
    return 'Welcome to Flask!!!'


@app.route('/ping')
def ping():
    return 'pong'


@app.route('/foo')
def foo():
    return '{"msg":"Hello from the foo microservice"}'


@app.route('/print/<name>')
def print(name=None):
    return '{{"msg":{} }}'.format(name)


"""
The entry point
Warning: Use 0.0.0.0 instead of 127.0.0.1 to avoid issue when use docker container   
"""
if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
