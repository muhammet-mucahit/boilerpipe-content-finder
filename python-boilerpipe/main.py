from typing import List

from boilerpipe.extract import Extractor
from fastapi import FastAPI
from pydantic import BaseModel

from multiprocessing import Pool
from multiprocessing import Process

from bs4 import BeautifulSoup

import requests
import re

app = FastAPI()

class JSONInput(BaseModel):
    urls: List[str] = []

class Boilerpipe():
    def __init__(self, url, result):
        self.url = url
        self.result = result

class BoilerpipeResult():
    def __init__(self, title, description, h1, content, wordCount):
        self.title = title
        self.description = description
        self.h1 = h1
        self.content = content
        self.wordCount = wordCount

def count_word_content(extractor):
    return len(extractor.getText().split())

def findInPageWithSelector(extractor, selector):
    item = BeautifulSoup(extractor.data,"html.parser").find(selector)
    if item:
        return item.text.strip()
    return ""

def findMetaInfos(extractor):
    beautifulSoup = BeautifulSoup(extractor.data,"html.parser")
    desc = beautifulSoup.find_all(attrs={"name": re.compile(r'Description', re.I)})
    if desc:
        return desc[0]['content']
    return ""

def get_boilerpipe(url):
    extractor = Extractor(extractor='ArticleExtractor', url=url)
    title = findInPageWithSelector(extractor, "title")
    h1 = findInPageWithSelector(extractor, "h1")
    description = findMetaInfos(extractor)
    content = extractor.getText()
    wordCount = count_word_content(extractor)
    
    result = BoilerpipeResult(str(title), str(description), str(h1), str(content), wordCount)
    boilerpipe = Boilerpipe(url, result)
    return boilerpipe

@app.post("/api/v1/boilerpipe")
def boilerpipe(json: JSONInput):
    result = map(get_boilerpipe, json.urls)
    return list(result)