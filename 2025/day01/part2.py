dial = 50
zeroes = 0

with open("input.txt") as f:
    for line in f:
        line = line.strip()
        direction = line[0]
        amount = int(line[1:])

        # Determine which way the dial moves per click
        delta = 1 if direction == "R" else -1

        # Simulate each click
        for _ in range(amount):
            dial = (dial + delta) % 100
            if dial == 0:
                zeroes += 1

print("Password:", zeroes)
