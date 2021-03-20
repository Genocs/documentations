from flask import Flask, request, jsonify

#from flask import render_template

app = Flask(__name__)


@app.route('/')
def hello():
    return 'Welcome to Flask!!!'


@app.route('/ping', methods=['GET'])
def ping():
    return 'pong'


@app.route('/order/submit', methods=['POST'])
def order_submit():
    content = request.json
    content["data"]["status"] = "Accepted"
    return jsonify(content)


@app.route('/print/<name>')
def print(name=None):
    return '{{"msg":{} }}'.format(name)


"""
The entry point
Warning: Use 0.0.0.0 instead of 127.0.0.1 to avoid issue when use docker container   
"""
if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5400)
