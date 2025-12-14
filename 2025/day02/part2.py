import re

invalid_id = re.compile(r'^([1-9]\d*)\1+$')  # some sequence of digits repeated at least twice.
total = 0

with open("input.txt", "r") as f:
    for line in f:
        parts = line.strip().split(",")
        for part in parts:
            part = part.strip()

            start_str, end_str = part.split("-", 1)
            start_num = int(start_str.strip())
            end_num = int(end_str.strip())

            # Loop over all numbers in the range
            for n in range(start_num, end_num + 1):
                if invalid_id.fullmatch(str(n)):
                    total += n

print("Sum:", total)
