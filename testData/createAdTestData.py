import random
import json
from datetime import datetime, timedelta

# Define the conditions
conditions_template = {
    "ageStart": 0,
    "ageEnd": 100,
    "gender": ["M", "F"],
    "country": ["JP", "US", "UK", "TW"],
    "platform": ["iOS", "Android"]
}

data_sets = []
i = 1
for _ in range(1000):
    # Generate end year with 80% chance for 2025 and 20% chance for 2024
    end_year = 2024

    # Generate conditions with each field being optional
    conditions = {
        "gender": random.choice(conditions_template["gender"]) if random.random() < 0.5 else None,
        "country": random.sample(conditions_template["country"], k=random.randint(1, len(conditions_template["country"]))) if random.random() < 0.5 else None,
        "platform": random.sample(conditions_template["platform"], k=random.randint(1, len(conditions_template["platform"]))) if random.random() < 0.5 else None,
    }

    # If ageStart is present, ensure ageEnd is also present
    if random.random() < 0.5:
        conditions["ageStart"] = random.randint(0, 100)
        conditions["ageEnd"] = random.randint(conditions["ageStart"], 100)

    # Remove None values
    conditions = {k: v for k, v in conditions.items() if v is not None}

    # Generate a random startAt time within the year 2021
    startAt = datetime(2024, 1, 1) + timedelta(days=random.randint(0, 180), hours=random.randint(0, 23))
    endAt = datetime(2024, 1, 1) + timedelta(days=random.randint(0, 180), hours=random.randint(0, 23))
    # Create the test data

    test_data = {
        "title": f"test{i}",
        "startAt": startAt.isoformat() + "Z",
        "endAt": endAt.isoformat() +"Z",
        "conditions": conditions
    }

    data_sets.append(test_data)
    i+=1

# Convert the test data to JSON
json_data = json.dumps(data_sets, indent=4)

# Print the JSON data
print(json_data)
