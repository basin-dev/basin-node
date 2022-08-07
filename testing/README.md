# Testing Basin

Clone the `basin-node` repo:
```
git@github.com:basin-dev/basin-node.git
```

## Setting up the test Twitter API

Enter the `testing` directory:
```
cd testing
```

Ensure you have Python 3.9:
```
python3 --version
```

Create a venv `env`:
```
python3 -m venv env
```

Activate virtual environment:
```
source env/bin/activate
```

Install FastAPI & Uvicorn:
```
pip3 install fastapi "uvicorn[standard]"
```

Run the server with:
```
uvicorn main:app --reload
```

Learn how to use the API:
```
 http://127.0.0.1:8000/docs
```

## Setting up the Basin node

Open a second tab with `src` as your current directory.

### Add a new keystore file for the given DID

*Question: what are our instructions for determining your DID?*
*Question: what is a key store file?*
*Question: why are each of these necessary?*
*Question: what does it mean to be the node's default signer?*

Enter a did, a name for your private key, and a password after running:
```
go run . auth add
```

### Extract and print the info from your keystore file

You can see your keystore file info by running:
```
go run . auth {did} {pw}
```

### Delete the keystore for a given DID

You can remove the keystore for a DID by running:
```
go run . auth forget {did}
```

### Start the Basin node

Specify your DID and password by running:
```
go run . up --did={did} --pw={pw}
```

## Setting up the interactive CLI

Open a third tab with `src` as your current directory again

Attach to the Basin node with interactive CLI by running:
```
go run . attach --http=http://127.0.0.1:8555
```

## Interacting with node in producer mode

From within the Basin node with interactive CLI.

Register your first resource:
```
register basin://tydunn.com.twitter.followers -a ../testing/config/adapter.json -p ../testing/config/permissions.yaml -s ../testing/config/schema.json
```

Try reading it:
```
do read basin://tydunn.com.twitter.followers
```


## Interacting with node in consumer mode

Question: what are the steps for the consumer?