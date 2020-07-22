# ------------------------------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------------------------------

import time
import requests
import os

dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)

dapr_url = "http://localhost:{}/v1.0/bindings/influx".format(dapr_port)
n = 0.0
while True:
    n += 1.0
    payload = { 
        "data": {
            "measurement": "temp",
            "tags": "room=dorm,building=building-a",
            "values": "sensor=\"sensor X\",avg={},max={}".format(n, n*2)
            }, 
        "operation": "create" 
    }
    print(payload, flush=True)
    try:
        response = requests.post(dapr_url, json=payload)
        print(response, flush=True)

    except Exception as e:
        print(e, flush=True)

    time.sleep(1)