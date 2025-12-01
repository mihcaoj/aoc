dial = 50
zeroes = 0

with open("input.txt", "r") as f:
    for line in f:
        line = line.strip()
        direction = line[0]
        amount = int(line[1:])
        delta = -amount if direction == 'L' else amount
        dial = (dial + delta) % 100
        zeroes += (dial == 0)
print("Password:", zeroes)
