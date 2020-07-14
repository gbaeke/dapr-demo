# Pubsub demo


## Subscriber

Node application that uses express. Listens on port 3000 for:

- GET /dapr/subscribe: dapr runtime calls this to find the topics the app is interested in; in this case the app subscribes to sampleTopic because we return:
 
`
[{"topic": "sampleTopic","route": "/sampler"}]
`


- POST /sampler: dapr will post messages on topic sampleTopic to this handler; the app just logs this and returns 200 to indicate proper receipt

Start the subscriber with:

dapr run --app-id sub --app-port 3000 node app.js

**Note:** you will notice that the actual message uses the CloudEvents spec


## Publisher

Node app that uses request to publish a increasing counter to topice sampleTopic at a 5 seconds interval. The message is POSTed to the following URL: http://localhost:3500/v1.0/publish/sampleTopic

Start the publisher with:

dapr run --app-id pub node app.js

**Note:** instead of running the publisher, you can publish messages with the dapr extension in VS Code

## PubSub backend on local machine

Backend is Redis Streams, automatically installed and configured during **dapr init**