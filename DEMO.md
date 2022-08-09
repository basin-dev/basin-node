# Demo

## Authentication

Obtain 2 key method DIDs and their respective base58 private keys from [here](https://did.key.transmute.industries/generate/ed25519). Henceforth referred to as `<DID_1>`, `<DID_2>`, `<PK_1>`, and `<PK_2>`.

Add these DIDs to the local Basin keystore (from `/src`) by twice running
```bash
go run . auth add
```
and entering the corresponding DID/PK. This demo assumes both passwords are "password".

## Uvicorn

Start data server (from `testing/`)
```bash
testing % uvicorn main:app --reload
```

## Node 1

Start node (from `src/`)
```bash
go run . up --did <DID_1> --pw password --http http://localhost:8555
```

Attach in another terminal
```bash
go run . attach --http http://127.0.0.1:8555
```

Register from interactive
```bash
register basin://tydunn.com.twitter.followers -a ../testing/config/adapter.json -p ../testing/config/permissions.yaml -s ../testing/config/schema.json
```

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

Start node (from `src/`)
```bash
go run . up --did <DID_2> --pw password --http http://localhost:8556
```

Attach in another terminal
```bash
go run . attach --http http://127.0.0.1:8556
```

Try to read the resource, should be denied
```bash
do read basin://tydunn.com.twitter.followers
```

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