from operator import xor

def getInput(input_name):
    "returns the input of the file as a list of the lines"
    f = open(input_name, "r")
    lines = f.readlines()
    lines = lines[9:]

    input = []

    for i in range(len(lines)):

        #last line no /n
        if(i == len(lines)-1):
            line = lines[i][:len(lines[i]) - 1]
        #first - last-1 lines
        else:
            line = lines[i][:len(lines[i]) - 2]

        input.append(line)

    return input


def getValidPasswworts(input):
    "line comes in form: '1-3 a: abcde' first the range than the letter to be checked and then the password"
    valid_old = 0
    valid_new = 0

    for line in input:
        #find all the needed indexes
        dash = line.find('-')
        colon = line.find(':')

        #find all needed parameters
        start = int(line[0:dash])
        end = int(line[dash+1:colon-2])
        letter_to_check = line[colon-1:colon]
        password = line[colon+2:]

        if isValid_old(start, end, letter_to_check, password):
            valid_old += 1

        if isValid_new(start, end, letter_to_check, password):
            valid_new += 1

    return valid_old, valid_new

def isValid_old(at_least, at_most, letter_to_check, password):
    "checks whether the given letter appears within the given range in the password, returns True or False respectively"
    number_of_times_appeared = 0
    for letter in password:
        if letter == letter_to_check:
            number_of_times_appeared += 1

    if number_of_times_appeared >= at_least and number_of_times_appeared <= at_most:
        return True
    else:
        return False

def isValid_new(first_occurrence, second_occurrence, letter_to_check, password):
    "checks whether the given letter appears exactly once in the password, returns True or False respectively, there is no counting by zero"

    appears = 0

    if password[first_occurrence-1:first_occurrence] == letter_to_check:
        appears += 1
    if password[second_occurrence-1:second_occurrence] == letter_to_check:
        appears += 1

    if appears != 1:
        return False
    else:
        return True


if __name__ == '__main__':
    input = getInput("Input-Day2.rtf")

    #print(input)
    print(len(input))
    print(getValidPasswworts(input))