from flask import Flask,request
app = Flask(__name__)

@app.route('/')
def data():
    gpio = request.args.get('gpio')
    duration = request.args.get('duration')
    print('gpio: %s, duration: %s')

