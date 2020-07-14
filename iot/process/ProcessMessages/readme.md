# Read from IoT Hub with Azure Functions and send alerts via Dapr PubSub

## Getting ready

Getting started with dapr: https://dapr.io/

See the dapr Azure Functions quickstart: https://github.com/dapr/azure-functions-extension/blob/master/docs/quickstart.md

You will need:

- dapr installed and configured locally (dapr init)
- Azure Functions Core Tools v3
- Docker
- .NET Core SDK
- Node

I have this all running inside WSL 2 with Docker in WSL 2 as well.

## Create the Function App

Created a new folder called **process** and ran

func init --docker

Created a new function called **ProcessMessages** using JavaScript and the IoT Hub trigger. This actually uses the Event Hub trigger.

## Install the Dapr extension and EventHubs

Remove **extensionBundle** from host.json

Run the following command:

func extensions install -p Dapr.AzureFunctions.Extension -v 0.8.0-preview01

Modify **extensions.csproj** and add the following to ItemGroup:

`<PackageReference Include="Microsoft.Azure.WebJobs.Extensions.EventHubs" Version="4.1.1">`

Run the following command:

dotnet build extensions.csproj -o bin --no-incremental

## Update the bindings in function.json

Update the **eventhubTrigger** binding with the Event Hub name that's used by your IoT Hub. In addition, add a reference to the connection string in **settings.local**.

Add a dapr binding to send PubSub messages to a topic called **alert**.

**Note**: on the local machine, dapr will use Redis Streams by default

## Update settings.local

Add the connection string to the Event Hub.

Add the connection string to the Azure Storage Account (required). You could also use the Azure Storage emulator for this (not tried).

## Update the code in index.js

Check the temperature value in the received JSON and use dapr PubSub to send the alert.

## Run the Azure Function locally

**DO NOT** press F5 to run the Azure Function locally to check if it works. The Azure Function needs to be run with dapr:

dapr run --app-id sub --port 3501 -- func start -p 7071

## Send a test message to the IoT Hub

Use the IoT Hub extension in Visual Studio code. Make sure an alert is triggered.

Check the alert message in redis:

xread STREAMS alert 0





