# birthyear, issue year, expiration year, height, eye color, passport ID, hair color
# with their valid inputs: always the min and the max of each category, height first 2 if cm second if in, eye color, musst be exactly one, hair color # followed by 6 numbers or letters a-f
# cid = country ID can be ignored

REQUIRED_FIELDS = [["byr:", (1920, 2002)], ["iyr:", (2010, 2020)], ["eyr:", (2020, 2030)], ["hgt:", (150, 193, 59, 76)],
                   ["ecl:", ("amb", "blu", "brn", "gry", "grn", "hzl", "oth")], ["pid:", (000000000, 999999999)],
                   ["hcl:", ('0', '9', "a", "f")]]


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
        if (i == len(lines) - 1):
            line = lines[i][:len(lines[i]) - 1]
        # first - last-1 lines
        else:
            line = lines[i][:len(lines[i]) - 2]

        input_list.append(line)

    return input_list


def howManyValidPassports(input):
    """
    makes each passport into a long string and checks wether its a valid passport, if yes than it counts one up

    :param input: list of passports, new passport is started by a line being empty in between
    :return: returns the count of valid passports in a given list
    """
    counter = 0
    currentPassport = ""
    for line in input:
        if line == "":
            if isValid(currentPassport):
                counter += 1

            currentPassport = line
            continue

        currentPassport += (line + " ")

        if line == input[len(input) - 1]:
            if isValid(currentPassport):
                counter += 1

    return counter


def hasValidInputs(passport):
    """
    checks wether or not the given Data in the passport String matches the requirements in the REQUIRED_FIELDS

    :param passport: a string that represents the passport
    :return: True or False respectively
    """
    for field in REQUIRED_FIELDS:
        indx = passport.find(field[0]) + 4

        # muss in einer bestimmten range sein
        if field[0] == "byr:" or field[0] == "iyr:" or field[0] == "eyr:" or field[0] == "pid:" or field[0] == "hgt:":
            number = passport[indx:]
            space = number.find(" ")
            number = number[:space]
            smallest = field[1][0]
            biggest = field[1][1]

            # chose correct comparison for hgt
            if "cm" in number:
                number = number[:len(number) - 2]
            if "in" in number:
                number = number[:len(number) - 2]
                smallest = field[1][2]
                biggest = field[1][3]

            # schauen ob pid 9 stellen hat
            if field[0] == "pid:":
                if 9 != len(number):
                    return False

            number = int(number)

            if number < smallest or number > biggest:
                return False
            continue

        # muss eine bestimmte Farbe bzw Bezeichnung sein
        if field[0] == "ecl:":
            ecl = passport[indx:]
            space = ecl.find(" ")
            ecl = ecl[:space]

            if ecl not in field[1]:
                return False
            continue

        #Farbe muss ein # haben und 6 Character entweder 0-9 oder a-f
        elif field[0] == "hcl:":
            hcl = passport[indx:]
            space = hcl.find(" ")
            hcl = hcl[:space]

            if hcl[0] != "#" or len(hcl) != 7:
                return False
            for c in hcl[1:]:
                if not (field[1][0] <= c <= field[1][1]) and not field[1][2] <= c <= field[1][3]:
                    return False
            continue

    return True


def isValid(passport):
    """
    looks wether the required fields are in the passport String and wether or not the fields have the required input
    """
    for field in REQUIRED_FIELDS:
        if field[0] not in passport:
            return False

    return hasValidInputs(passport)


if __name__ == '__main__':
    input = getInput("Input-Day4.rtf")
    test = getInput("Test-4.rtf")
    invalid = getInput("invalid-Day4.rtf")
    valid = getInput("Valid-Day4.rtf")
    print(howManyValidPassports(valid))
    print(howManyValidPassports(input))
