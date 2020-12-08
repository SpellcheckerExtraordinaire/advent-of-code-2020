def getInput(input_name):
    "returns the input of the file as a list of the lines"
    f = open(input_name, "r")
    lines = f.readlines()
    lines = lines[9:]

    input = []

    for i in range(len(lines)):

        if(i == len(lines)-1):
            line = lines[i][:len(lines[i]) - 1]
        else:
            line = lines[i][:len(lines[i]) - 2]
        #print(line)
        input.append(int(line))

    return input

def pruneList(input, min):
    "prunes list of numbers that together with the min will add up to be greater than 2020, returns a new list that only aacepts"
    print(min)
    l = []
    for number in input:
        x = number + min
        if (x) <= 2020:
            l.append(number)

    return l

def find2020_2(input):
    "finds two numbers that add up to 2020"
    input.sort()
    while(True):
        min = input[0]
        input = pruneList(input, min)

        for number in input:
            if number == min:
                continue
            if number + min == 2020:

                return number, min

        input.remove(min)


def find2020_3(input):
    "finds three numbers that add up to 2020"

    input.sort()
    a = input[0]
    b = input[1]
    input = pruneList(input, a + b)

    for i in range(0, len(input)):
        a = input[i]
        for j in range(i+1, len(input)):
            b = input[j]
            for k in range(j+1, len(input)):
                c = input[k]

                if a + b + c == 2020:
                    return a, b, c


if __name__ == '__main__':
    #print(getInput())
    input = getInput("Input-Day1.rtf")
    min = min(input)

    #numbers = find2020_2(input)
    #print(numbers[0]*numbers[1])
   # print(find2020_3(input))

    test = getInput("test-day1.rtf")
    print(find2020_3(test))
    numbers = find2020_3(input)
    print(numbers[0] * numbers[1] * numbers[2])





