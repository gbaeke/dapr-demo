const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());


const port = 3000;

// mqtt component will post messages from influx topic here
app.post('/mqtt', (req, res) => {
    console.log("MQTT Binding Trigger");
    console.log(req.body)

    // body is expected to contain room and temperature
    room = req.body.room
    temperature = req.body.temperature

    // room should not contain spaces
    room = room.split(" ").join("_")

    // create message for influx component and signalr component
    message = {
        "measurement": "stat",
        "tags": `room=${room}`,
        "values": `temperature=${temperature}`,
        "target": "newMessage", 
        "arguments": [req.body]
    };
    
    // send the message to influx output binding
    res.send({€
        "to": ["influx", "signalr"],
        "concurrency": "parallel",
        "data": message
    });
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));