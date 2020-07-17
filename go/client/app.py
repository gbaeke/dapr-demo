import time
import requests
import os

dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)
dapr_url = "http://localhost:{}/v1.0/invoke/goserver/method/HelloFromGo".format(dapr_port)

print("Let's start...")

n = 0
while True:
    n += 1
    print("Create a message...")
    message = {"messageId": n}
    
    try:
        response = requests.post(dapr_url, json=message, timeout=5)
        print(response.text)
    except Exception as e:
        print(e, flush=True)

    time.sleep(1)
