import heapq
import math

throws = []
items = []
operator = []
value = []
divisor = []
targets = []

lcd = 1

def solve():
    monkeys = range(len(throws))
    for round in range(10000):
        for monkey in monkeys:
            handle_monkey(monkey)

    top_monkeys = heapq.nlargest(2, throws)
    return top_monkeys[0] * top_monkeys[1]

def handle_monkey(index):
    for item in items[index]:
        v = value[index]
        if v == 'old':
            if operator[index] == '+':
                worry_level = (item + item) % lcd
            else:
                worry_level = (item * item) % lcd
        else:
            if operator[index] == '+':
                worry_level = (item + v) % lcd
            else:
                worry_level = (item * v) % lcd
        #worry_level = math.floor(worry_level/3)

        if worry_level % divisor[index] == 0:
            items[targets[index][0]].append(worry_level)
        else:
            items[targets[index][1]].append(worry_level)
        throws[index] += 1
    items[index].clear()


def main():
    with open('input.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()
    
    global lcd

    i = 0
    while( i < len(input)):
        throws.append(0)
        curr_items = input[i + 1]
        curr_items = curr_items.split()
        curr_items = curr_items[2:]
        curr_items = [int(x.strip(',')) for x in curr_items]

        items.append(curr_items)

        curr_op = input[i + 2]
        curr_op = curr_op.split()
        operator.append(curr_op[4])
        if curr_op[5] == 'old':
            value.append('old')
        else:
            value.append(int(curr_op[5]))
        
        curr_div = input[i + 3]
        curr_div = curr_div.split()
        lcd *= int(curr_div[-1])
        divisor.append(int(curr_div[-1]))

        curr_target_t = input[i + 4]
        curr_target_t = curr_target_t.split()[-1]

        curr_target_f = input[i + 5]
        curr_target_f = curr_target_f.split()[-1]

        targets.append((int(curr_target_t), int(curr_target_f)))
        i += 7

    print(solve())
    
    print(throws)

if __name__ == "__main__":
    main()