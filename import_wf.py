import sys, os
import pymongo
import bson
import re
from pymongo import MongoClient
from os import listdir
from os.path import isfile, join


def main(argv):
  client = MongoClient()
  db = client.lerp
  SCRIPT_DIR = argv[0] if len(argv)>0 else 'wf/'

  for f in listdir(SCRIPT_DIR):
    if isfile(join(SCRIPT_DIR,f)) and f.startswith("wf-") and f.endswith(".js"):
    	wfId, _ = f.split('.', 1)
    	wfId = wfId.replace("-", "_")
    	file_path = os.path.join(SCRIPT_DIR, f)
    	
    	with open(file_path, 'r') as content_file:
    		content = content_file.read()
	    	code = bson.code.Code(content)
	    	db.system.js.save({ '_id': wfId, 'value': code })   
	    	print(wfId + " workflow loaded.")
    

if __name__ == "__main__":
   main(sys.argv[1:])