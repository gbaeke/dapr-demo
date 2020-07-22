# InfluxDB 2.0 Binding

## Run InfluxDB 2 container

docker run -p 9999:9999 -v $PWD:/var/lib/influxdb quay.io/influxdb/influxdb:2.0.0-beta

Above command will run InfluxDB 2 on port 9999 and use the local directory for state.

## Download influx client

wget https://dl.influxdata.com/influxdb/releases/influxdb_client_2.0.0-beta.14_linux_amd64.tar.gz

tar xvfz influxdb_client_2.0.0-beta.14_linux_amd64.tar.gz

Now copy influx to your path

## Run influx setup

influx setup

Enter primary username

Enter password

Primary organization name

Primary bucker name

Retention period (in hours)

Config stored in $HOME/.influxdbv2/configs

## List buckets

influx bucket list

Copy the id of the bucket you want to work with

## Create auth token

influx auth create -o Inity --write-bucket BUCKETID --read-bucket BUCKETID

## Writing data - influx-go/main.go

Uses the line API to write data:

line:="measurement,tagName=tagValue val1=x,val2=y

Remember:
- data is saved in a bucket (bucket passed in component spec); a bucket belongs to an organization (organization passed in component spec)
- measurement: container for tags, fields and timestamps
- fields: a field key & a key value; field values can be strings, floats, integers or booleans; *Note* that fields are *NOT* indexed
- tags: tag keys and values stored as strings and metadata; tags are *optional* and tags are *indexed*
- series key: collection of points that *share* a measurement, tag set and field key
- series: timestamps and field values for a given series key
- point: series key + fields value + timestamp

See https://v2.docs.influxdata.com/v2.0/reference/key-concepts/data-elements/ for more information