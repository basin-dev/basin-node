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
