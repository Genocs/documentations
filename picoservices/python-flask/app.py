from flask import Flask, request, jsonify

#from flask import render_template

app = Flask(__name__)


@app.route('/')
def hello():
    return 'Welcome to Flask!!!'


@app.route('/ping', methods=['GET'])
def ping():
    return 'pong'


@app.route('/post_my_data', methods=['POST'])
def post_my_data():
    content = request.json
    return jsonify(content)


@app.route('/print/<name>')
def print(name=None):
    return '{{"msg":{} }}'.format(name)


"""
The entry point
Warning: Use 0.0.0.0 instead of 127.0.0.1 to avoid issue when use docker container   
"""
if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
