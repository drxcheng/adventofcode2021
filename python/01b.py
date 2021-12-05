inputFile = open('01.txt', 'r')
lines = inputFile.readlines()

depths = []
for line in lines:
    depths.append(int(line.strip()))

numberOfIncreases = 0
prevSlidingWindowValue = depths[0] + depths[1] + depths[2]
index = 3
while index < len(depths):
    slidingWindowValue = depths[index - 2] + depths[index - 1] + depths[index]
    if slidingWindowValue > prevSlidingWindowValue:
        numberOfIncreases += 1
    prevSlidingWindowValue = slidingWindowValue
    index += 1

print(numberOfIncreases)
