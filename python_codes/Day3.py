# konstanten
TREE = '#'
FREE = '.'
# slopes in format right, down
SLOPES = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]


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


def howManyTrees(input, x_slope, y_slope):
    """
    :param input: List of Strings that represent the hill. # = Tree, . = open spot, the input repeats itself to the right indefinetly
    :param x_slope: how many steps one takes to the right
    :param y_slope: how many steps one takes to the left
    :return: number of trees one encounters when traversing the hill (input) with the given slope from the left top corner torwards the right bottom
    """
    number_of_trees = 0
    x = 0
    for line in input[::y_slope]:
        if line[x] == TREE:
            number_of_trees += 1

        x = (x + x_slope) % len(line)

    return number_of_trees


def allSlopes(input):
    """
    checks how many trees one encounters for each slope in SLOPES and returns the product

    :param input: List of Strings that represent the hill. # = Tree, . = open spot, the input repeats itself to the right indefinetly
    :return: product of all Tree encounters for each slope
    """
    product = 1
    for slope in SLOPES:
        product *= howManyTrees(input, slope[0], slope[1])

    return product


if __name__ == '__main__':
    test = getInput("Test-3.rtf")
    input = getInput("Input-Day3.rtf")
    # print(test)
    # print(howManyTrees(input))
    print(allSlopes(input))
