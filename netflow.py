import csv
import json
from collections import defaultdict

def read_and_display_top10csv(filename):
    activity = defaultdict(int)
    
    with open(filename, "r") as file:
        readcsv = csv.DictReader(file)
        
        for row in readcsv:
            sourceip = row['source_ip']
            packets = int(row['packets'])

            activity[sourceip] += packets

        activity = sorted(activity.items(), key= lambda x : x[1], reverse=True)

    print(f"total number of logs: {len(activity)}")
    print(f"top 10 activity:")
    for ip, packs in activity[:10]:
        print(f"ip: {ip} and packtes: {packs}")
    



def read_and_display_top10json(filename):
    activity = defaultdict(int)
    
    with open(filename, "r") as file:
        readjson = json.load(file)
        
        for entry in readjson:
            sourceip = entry['source_ip']
            packets = entry['packets']

            activity[sourceip] += packets

        activity = sorted(activity.items(), key= lambda x : x[1], reverse=True)

    print(f"total number of logs: {len(activity)}")
    print(f"top 10 activity:")
    for ip, packs in activity[:10]:
        print(f"ip: {ip} and packtes: {packs}")
    


read_and_display_top10json('netflow.json')