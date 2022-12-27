

def assignment_1(cycle, counter, sum):
    if (cycle <= 220) and (cycle % 40 == 20):
                sum += counter * cycle
    return sum

def draw_pixel(crt_image, cycle, counter):
    print((cycle - 1) % 40, counter)
    pixel = cycle - 1

    if pixel % 40 in [counter - 1, counter, counter + 1]:
        crt_image = crt_image + '#' 
    else:
        crt_image = crt_image + '.'

    if cycle % 40 == 0:
        crt_image = crt_image + '\n'

    return crt_image

def main():
    with open('input_day10.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()
    

    counter = 1
    # sum = 0
    cycle = 1

    crt_image = ''

    for line in input:
        split = line.split()

        oper = split[0]

        crt_image = draw_pixel(crt_image, cycle, counter)

        if oper == 'noop':
            cycle += 1
            continue
        else:
            value = int(split[1])
            cycle += 1
            crt_image = draw_pixel(crt_image, cycle, counter)
            counter += value
            cycle += 1

    print(crt_image)




if __name__ == "__main__":
    main()