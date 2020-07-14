# Azure Functions and DAPR

## HttpTrigger function

Create a new function with Core Tools v3:

- func init --docker
- func new

Install Dapr extension:

- remove extensionBundle from host.json
- run func extensions install -p Dapr.AzureFunctions.Extension -v 0.8.0-preview01

Run the function with Dapr

dapr run --app-id function-app --port 3501 -- func start -p 7071

Now POST to the function

curl -XPOST -H "Content-type: application/json" -d 'Hello' 'http://localhost:7071/api/state/mykey'

The above will do the following:

- append Hello to key mykey in the state store (by default in Redis)
- to verify in Redis: from redis-cli, run hgetall function-app||mykey
- above, function-app is the dapr app id and mykey the key

## Topic subscriber function

- func new -> use Timer Trigger (will be changed later to Dapr Topic Trigger)
- name of function: DaprSubscribeTrigger
- see function.json and index.js

Now run the Azure Function (kill previously running one):

dapr run --app-id function-app --app-port 3001 --port 3501 -- func start -p 7071

**Note:** the subscriber needs to receive trigger events FROM the sidecar; sidecar delivers them on port 3001 (app-port above)
