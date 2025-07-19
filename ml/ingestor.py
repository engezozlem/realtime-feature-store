import csv
import redis

r = redis.Redis(host="localhost", port=6379, decode_responses=True)

with open("./sample_users.csv", newline="") as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        entity_id = row.pop("entity_id")
        r.hset(entity_id, mapping=row)
        print(f"Ingested features for entity_id={entity_id}: {row}")
