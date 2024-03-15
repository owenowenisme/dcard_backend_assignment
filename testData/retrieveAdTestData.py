import random
import json

genders = ["M", "F", ""]
countries = ["JP", "US", "UK", "TW", ""]
platforms = ["iOS", "Android", ""]

data = []
for _ in range(10):
    offset = str(random.randint(0, 50)) if random.random() < 0.8 else ""
    limit = str(random.randint(0, 50)) if random.random() < 0.8 else ""
    age = str(random.randint(10, 50)) if random.random() < 0.8 else ""
    gender = random.choice(genders)
    country = random.choice(countries)
    platform = random.choice(platforms)
    data.append([offset, limit, age, gender, country, platform])

with open('test_data.json', 'w') as f:
    json.dump(data, f, indent=4)
