END = "no other bags"
BAG = "shiny gold"


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


def makeRuleList(input_list):
    """
    takes each sentence and creates a list of the important details. sentences are structured as follows:
    pale turquoise bags contain 3 muted cyan bags, 5 striped teal bags.
    light maroon bags contain no other bags.

    :param input_list: list of the lines in the input
    :return: list of lists, list of all rules, each rule separated as follows ["name", "color_contains", number_it_contains, ...]
    """
    l = []
    for line in input_list:
        rule = []
        bag = line.find("bag")
        rule_name = line[:bag - 1]

        rule.append(rule_name)

        if END in line:
            rule.append(END)
            l.append(rule)
            continue

        contain = line.find("contain")
        line = line[contain + len("contain") + 1:]
        number = line

        # get number and bag_name that have commas behind them
        while line.count(",") != 0:
            # get number
            space = number.find(" ")
            bag_name = number[space + 1:]
            number = number[:space]

            # get bag_name
            comma = line.find(",")
            line = bag_name[comma:]
            bag = bag_name.find("bag")
            bag_name = bag_name[:bag - 1]

            rule.append(bag_name)
            rule.append(int(number))
            number = line

        # get last bag_name and number
        space = number.find(" ")
        bag = number.find("bag")
        bag_name = number[space + 1:bag]
        number = number[:space]
        bag_name = bag_name[:-1]

        rule.append(bag_name)
        rule.append(int(number))

        l.append(rule)

    return l


def calculateNumberOfBags(rules):
    """

    :param rules: list of lists
    :return: tuple, (number of bags that can eventually contain at least one BAG, number of bags hat a BAG contains)
    """
    counter = 0
    bags_within_BAG = 0
    for rule in rules:
        # we want to know in which bags we can put our wanted bag that's why we skip it
        if rule[0] == BAG:
            bags_within_BAG = bagsWithinBAG(rules, rule)
            continue
        counter += containsWantedBag(rules, rule)

    return counter, bags_within_BAG


def containsWantedBag(rules, rule):
    """

    :param rules: list of lists, is given through and not changed
    :param rule: list, current rule that is looked at
    :return: 1 if the outermost bag can contain a BAG
    """
    if rule[1] == END:
        return 0

    for line in rule:
        if line == BAG:
            return 1

    counter = 0
    for bag in rule[1::2]:
        current_bag = bag
        rule = findRule(rules, current_bag)
        counter += containsWantedBag(rules, rule)
        # we dont care if the bag could contain our wanted bag multiple times, at least once is sufficient
        if counter != 0:
            break

    return counter


def bagsWithinBAG(rules, rule):
    """

    :param rules: list of lists, is given through and not changed
    :param rule: list, current rule that is looked at
    :return: number of bags the through the rule specified bag contains
    """
    if rule[1] == END:
        return 0

    counter = 0
    #looking through the bags with index ; index+1 are the numbers
    for index in range(1, len(rule), 2):
        current_bag = rule[index]
        new_rule = findRule(rules, current_bag)
        counter += rule[index + 1]
        counter += rule[index + 1] * bagsWithinBAG(rules, new_rule)

    return counter


def findRule(rules, current_bag):
    """

    :param rules: list of lists, is given through and not changed
    :param current_bag: string, with name of a bag that we want to find
    :return: list of rules for that bag
    """
    for rule in rules:
        if rule[0] == current_bag:
            return rule


if __name__ == '__main__':
    test = getInput("test-Day7.rtf")
    test_rules = makeRuleList(test)
    input_rules = makeRuleList(getInput("Input-Day7.rtf"))
    print(calculateNumberOfBags(input_rules))
