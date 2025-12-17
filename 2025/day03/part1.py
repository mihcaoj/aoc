def max_joltage(bank: str) -> int:
    digits = list(map(int, bank.strip()))
    n = len(digits)

    # suffix_max[i] = largest digit to the right of index i
    suffix_max = [0] * n
    suffix_max[-1] = -1

    # Build suffix max array from right to left
    for i in range(n - 2, -1, -1):
        suffix_max[i] = max(suffix_max[i + 1], digits[i + 1])

    best = 0
    for i in range(n - 1):
        candidate = 10 * digits[i] + suffix_max[i]
        best = max(best, candidate)
    return best


total_joltage = 0

with open("input.txt", "r") as f:
    for line in f:
        bank_joltage = max_joltage(line)
        total_joltage += bank_joltage

print("Total output joltage:", total_joltage)
