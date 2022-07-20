Create a venv `env` using Python 3.9:
```
python3 -m venv env
```

Activate virtual environment:
```
source env/bin/activate
```

Install FastAPI:
```
pip3 install fastapi
```

Install Uvicorn:
```
pip3 install "uvicorn[standard]"
```

Run the server with:
```
uvicorn main:app --reload
```

Learn how to use the API:
```
 http://127.0.0.1:8000/docs
 ``