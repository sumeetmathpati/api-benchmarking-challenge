import os
import time
from datetime import datetime
import uuid
from fastapi import FastAPI
from elasticsearch7 import AsyncElasticsearch

app = FastAPI()
es = AsyncElasticsearch(hosts=os.getenv("ES_HOST", "localhost"))

INDEX = "events"


@app.get("/")
async def events():
    await es.create(
        index=INDEX,
        id=str(uuid.uuid4()),
        document={"timestamp": datetime.now()},
        refresh=False
    )
    resp = await es.count(index=INDEX)
    return {"event_count": resp["count"]}
