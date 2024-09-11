

def move_right(pos):
    return (pos[0] + 1, pos[1])

def move_left(pos):
    return (pos[0] - 1, pos[1])

def move_up(pos):
    return (pos[0], pos[1] + 1)

def move_down(pos):
    return (pos[0], pos[1] - 1)

def move_tail(tail, head):

    dx = abs(tail[0] - head[0])
    dy = abs(tail[1] - head[1])

    next_x = tail[0]
    next_y = tail[1]

    # only row difference
    if dx > 1 and dy == 0:
        if tail[0] > head[0]:
            next_x = tail[0] - 1
        else:
            next_x = tail[0] + 1
    # only col difference
    elif dy > 1 and dx == 0:
        if tail[1] > head[1]:
            next_y = tail[1] - 1
        else:
            next_y = tail[1] + 1
    elif (dx > 1 and dy > 0) or (dx > 0 and dy > 1):
        # new row pos
        if tail[0] > head[0]:
            next_x = tail[0] - 1
        else:
            next_x = tail[0] + 1
        # new col pos
        if tail[1] > head[1]:
            next_y = tail[1] - 1
        else:
            next_y = tail[1] + 1

    return (next_x, next_y)



def main():
    with open('input_day9.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()

    tail_list = []

    for i in range(10):
        tail_list.append((0, 0))

    tail_visited = set()
    tail_visited.add((0, 0))

    for line in input:
        split = line.split()
        direction = split[0]
        reps = int(split[1])

        for _ in range(reps):
            if direction == 'R':
                tail_list[0] = move_right(tail_list[0])
            elif direction == 'L':
                tail_list[0] = move_left(tail_list[0])
            elif direction == 'U':
                tail_list[0] = move_up(tail_list[0])
            else:
                tail_list[0] = move_down(tail_list[0])
            
            for i in range(1, len(tail_list)):
                tail_list[i] = move_tail(tail_list[i], tail_list[i - 1])

            if tail_list[9] not in tail_visited:
                tail_visited.add(tail_list[9])
    
    print(len(tail_visited))
    
            

    


if __name__ == "__main__":
    main()