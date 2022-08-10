# Demo

## Authentication

Obtain 2 key method DIDs and their respective base58 private keys from [here](https://did.key.transmute.industries/generate/ed25519). Henceforth referred to as `<DID_1>`, `<DID_2>`, `<PK_1>`, and `<PK_2>`.

Add these DIDs to the local Basin keystore (from `/src`) by twice running
```bash
go run . auth add
```
and entering the corresponding DID/PK. This demo assumes both passwords are "password".

## Uvicorn

> Ty heard about Basin and wants to join the network as a producer. He decides to make his list of Twitter followers available on the network

> He starts the server that runs the API that serves this Twitter followers data (we imagine most folks, who have substantial data, will already have infrastructure like this that Basin can wrap around).

Curl the data with
```bash
curl http://localhost:8000/followers
```

Start data server (from `testing/`, make sure venv is started)
```bash
uvicorn main:app --reload
```

## Node 1

> He then grabs the Basin code from the open source repo (in the long run, we would like to make this easier e.g. by packaging it as a Docker container)

> The first step is to set himself up as a producer on the network is generate DIDs and private keys.

> He then needs to add this authentication. To keep the demo moving, I already did this before. So now, I’ll just log in with those credentials

> The next step is to use interactive command line interface (CLI) app to interact with the Basin node (in the future, we would like this to be a GUI, but while we are iterating on the prototype, a CLI is better for us)

Start node (from `src/`)
```bash
go run . up --did <DID_1> --pw password --http http://localhost:8555
```
>> And for now need to paste the public key into `everything.go`

Attach in another terminal
```bash
go run . attach --http http://127.0.0.1:8555
```

 > I attach myself to the Basin node and then register the API endpoint we saw earlier as a Basin resource, defining the adapter (explain this), schema (explain this), and initial permissions (explain this)

Register from interactive
```bash
register basin://tydunn.com.twitter.followers -a ../testing/config/adapter.json -p ../testing/config/permissions.yaml -s ../testing/config/schema.json
```

> Now let’s try reading it: [show how because you have permissions for your own Basin resources, you get the same response as when you did the curl command]

and test read
```bash
do read basin://tydunn.com.twitter.followers
```

## Web Interface

Start web interface (from `/basin-ui`)
```bash
npm run dev
```

Show around whatever is there

## Node 2

> So then if I come along and want to use Ty’s followers data, I do the same process of setting up a node but act as a consumer

Start node (from `src/`)
```bash
go run . up --did <DID_2> --pw password --http http://localhost:8556
```

Attach in another terminal
```bash
go run . attach --http http://127.0.0.1:8556
```

> I try reading the data. It doesn’t work because I have not been approved to use his data and given the necessary permissions

Try to read the resource, should be denied
```bash
do read basin://tydunn.com.twitter.followers
```

> I need to request a subscription to his basin resource: [send request]

Request read subscription
```bash
subscribe basin://tydunn.com.twitter.followers <DID_2> --action read
```

## Node 1

Show that the new request is shown (from the attached interactive terminal)
```bash
do read basin://<DID_1>.basin.producer.requests
```

Approve the request? Do we want to attach IDs to the requests? Maybe have `requests --list` and `requests approve <REQUEST_ID>` commands?

## Node 2

Try to read the resource again, this time it should work (from attached interactive terminal)

```bash
do read basin://tydunn.com.twitter.followers
```

## The End!