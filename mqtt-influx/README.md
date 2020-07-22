# MQTT to Influx

This application subscribes to an MQTT topic on a local Mosquitto server and forwards the data to a local Influx DB container.

## Start the Influx DB server

Run the following command:

docker run -p 9999:9999 -v $PWD:/var/lib/influxdb quay.io/influxdb/influxdb:2.0.0-beta

**Note:** run this from an empty folder; Influx DB will save the data there

Navigate to http://localhost:9999 for initial configuration (user, password, org and bucket)

## Start the MQTT server

docker run -it -p 1883:1883 -p 9001:9001 eclipse-mosquitto

If you want to specify a custom mosquitto.conf:

docker run -it -p 1883:1883 -p 9001:9001 -v mosquitto.conf:/mosquitto/config/mosquitto.conf eclipse-mosquitto

## Test the MQTT server

Use a tool like MQTT.fx to subscribe to a topic and send a message to it.

Check for connection of the client in the output:

1595416922: New connection from 172.17.0.1 on port 1883.

## Run the code with Dapr

Run **nmp install** first

Run the code:

dapr run --app-id mqqtinflux --app-port 3000 --components-path=./components node app.js

**Note:** the component expects MQTT messages on a topic called **influx**

Use MQTT.fx to send a message on topic **influx** with the following contents:

{ "room": "somerome", "temperature": 22.7 }

The app will pick this up and save it (measurement = stat)

## Tip: add a SignalR component and send to both Influx AND SignalR




