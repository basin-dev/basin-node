from fastapi import FastAPI
from fastapi.responses import FileResponse

app = FastAPI()

@app.get("/")
async def help():
    return "Learn how to use API: http://127.0.0.1:8000/docs"

@app.get("/followers")
async def read_followers():
    return FileResponse('data/followers.json')

@app.get("/followers2")
async def read_followers2():
    return FileResponse('data/followers.json')

@app.get("/following")
async def read_following():
    return FileResponse('data/following.json')