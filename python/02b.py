inputFile = open('02.txt', 'r')
lines = inputFile.readlines()

offsetHorizontal = 0
offsetVertical = 0
aim = 0
for line in lines:
    data = line.split()
    direction = data[0].strip()
    amount = int(data[1].strip())
    if direction == 'forward':
        offsetHorizontal += amount
        offsetVertical += amount * aim
    elif direction == 'down':
        aim += amount
    elif direction == 'up':
        aim -= amount
    else:
        print(f'wrong direction in line {line}')

print(f'horizontal position: {offsetHorizontal}')
print(f'vertical position: {offsetVertical}')
print(abs(offsetHorizontal * offsetVertical))
