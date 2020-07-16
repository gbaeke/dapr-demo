const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());


const port = 3000;

// cron component calls this endpoint
app.post('/mqtts', (req, res) => {
    console.log("MQTT Binding Trigger");
    console.log(req.body)
   

    // client app listens for tarhet newMessage
    message = {
        "target": "newMessage", 
        "arguments": [req.body]
    };
    
    
    // send the message to signalr output binding
    res.send({
        "to": ["signalr"],
        'data': message
    });
});

app.listen(port, () => console.log(`Node App listening on port ${port}!`));