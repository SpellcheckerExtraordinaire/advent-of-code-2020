def getInput(input_name):
    """returns the input of the file as a list of the lines
    works perfect for .rtf files
    .txt files need adjusting"""
    f = open(input_name, "r")
    lines = f.readlines()
    lines = lines[9:]

    input_list = []

    for i in range(len(lines)):

        # last line no /n
        if i == len(lines) - 1:
            line = lines[i][:len(lines[i]) - 1]
        # first - last-1 lines
        else:
            line = lines[i][:len(lines[i]) - 2]

        input_list.append(line)

    return input_list


def getSum(lines):
    """
    takes the list of lines and puts one group, separated by '' into one current_line, then checks if letters in curent_line appears as many times as we have members in that group, takes the sum of all answers

    :param lines: list of lines, groups separated through ''
    :return:int, sum of all answers of how many times a letter a-z appears in a group exactly as many times as members of that group
    """
    current_line = ""
    counter = 0
    members = 0

    for line in lines:
        if line == '':
            counter += numberOfNewLetters(current_line, members)
            current_line = ""
            members = 0
            continue

        current_line += line
        members += 1

    counter += numberOfNewLetters(current_line, members)
    return counter


def numberOfNewLetters(line, members):
    """

    :param line: string of letters, answers the group answered yes to
    :param members: int, number of people in that group
    :return: int, how many times all members of the group answered yes to the same question
    """
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    counter = 0

    for letter in alphabet:
        if line.count(letter) == members:
            counter += 1

    return counter


if __name__ == '__main__':
    input = getInput("Input-Day6.rtf")
    print(input)
    print(getSum(input))
