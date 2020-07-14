// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

const express = require('express');
const bodyParser = require('body-parser');

const app = express();

// When Dapr sends a message it uses application/cloudevents+json
app.use(bodyParser.json({ type: 'application/*+json' }));

const port = 3000;

// tell Dapr what we want to subscribe to
// Dapr does a GET to /dapr/subscribe to find out
app.get('/dapr/subscribe', (_req, res) => {
    res.json([
        {
            "topic": "sampleTopic",
            "route": "/sampler"
        }

    ]);
});

// Dapr will post to the route configured for the topic
app.post('/sampler', (req, res) => {
    console.log("message on SampleTopic: ", req.body);

    // send 200 to indicate succesful receipt
    res.sendStatus(200);
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));