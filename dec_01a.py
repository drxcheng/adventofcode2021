inputFile = open('input.txt', 'r')
lines = inputFile.readlines()

numberOfIncreases = 0
prevDepth = None
for line in lines:
    depth = int(line.strip())
    if prevDepth is not None and depth > prevDepth:
        numberOfIncreases += 1
    prevDepth = depth

print(numberOfIncreases)
