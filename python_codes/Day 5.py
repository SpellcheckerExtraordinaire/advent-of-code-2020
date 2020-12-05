def getInput(input_name):
    """returns the input of the file as a list of the lines"""
    f = open(input_name, "r")
    lines = f.readlines()
    lines = lines[9:]

    input = []

    for i in range(len(lines)):

        # last line no /n
        if (i == len(lines) - 1):
            line = lines[i][:len(lines[i]) - 1]
        # first - last-1 lines
        else:
            line = lines[i][:len(lines[i]) - 2]

        input.append(line)

    return input

def getRowOrColumn(boarding_pass, low, high, lower, upper):
    """
    gets you the row or column that is specified, F or L Take the lower half of the looked at range B or R the upper half, starting with the range 0-127 or 0-7
    :param boarding_pass: 10 Letter String, first 7  F or B last 3 L or R
    :param low: lowest number in range, 0
    :param high: highest number in range, either 127 or 7
    :param lower: letter that keeps the lower half
    :param upper: letter that keeps the higher half
    :return: the row or coloumn that the pass specifies
    """
    for letter in boarding_pass:
        if letter == lower:
            range = (high-low)//2
            high = low + range
        if letter == upper:
            range = (high - low) // 2
            low = high - range

    return low

def getSeatIDs(input):
    """
    takes each element in the input list and gets the row and coloumn and than calculates the specific seatID

    :param input: list of all boarding passes
    :return: sorted list of all seating Ids
    """
    l = []

    for boarding_pass in input:
        row = getRowOrColumn(boarding_pass, 0, 127, 'F', 'B')
        column = getRowOrColumn(boarding_pass, 0, 7, 'L', 'R')
        l.append(row*8 + column)

    l.sort()
    return l


def findMySeatID(seatIds):
    """
    finds the missing seat where the +1 and -1 neighbours exist, not checking the first and last row
    :param seatIds: sorted list of seat Ids
    :return: the missing seat in the list of seats
    """

    for i in range(8, len(seatIds)-7):
        before = seatIds[i-1]
        current = seatIds[i]
        after = seatIds[i+1]
        if before != (current - 1) and current != (after + 1):
            return seatIds[i] - 1


if __name__ == '__main__':
    test = getInput("Test-5.rtf")
    input = getInput("Input-Day5.rtf")
    #print(test)
    ids = getSeatIDs(input)
    print(findMySeatID(ids))

