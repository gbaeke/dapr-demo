module.exports = function (context, IoTHubMessages) {
    context.log(`JavaScript eventhub trigger function called for message array: ${IoTHubMessages}`);
    
    IoTHubMessages.forEach(message => {
        // log the full message
        context.log(`Processed message: ${JSON.stringify(message)}`);

        // alert if temperature is high
        if (message.temperature > 29.5) {
            context.log('Sending alert message...')
            context.bindings.alert = 
            {
                "payload": "I don't care about the actual value. It's too hot!!!!",
                "topic": "alert"
            }
        }
    });

    
    

    context.done();
};