

def visible(start_row, start_col, graph, rows, cols):
    
    value = graph[start_row][start_col]
    scenic_score = 1
    #left
    trees = 0
    for i in range(start_col - 1, -1, -1):
        trees += 1
        if graph[start_row][i] < value:
            continue
        else:
            break
    
    scenic_score *= trees
    trees = 0
    #down
    for i in range(start_row + 1, rows):
        trees += 1
        if graph[i][start_col] < value:
            continue
        else:
            break

    scenic_score *= trees
    trees = 0
    #up
    for i in range(start_row - 1, -1, -1):
        trees += 1
        if graph[i][start_col] < value:
            continue
        else:
            break

    scenic_score *= trees
    trees = 0
    #right
    for i in range(start_col + 1, cols):
        trees += 1
        if graph[start_row][i] < value:
            continue
        else:
            break
    scenic_score *= trees

    return scenic_score


def main():
    
    with open('input_day8.txt', 'r', encoding='utf-8') as infile:
        input = infile.readlines()

    graph = []
    for line in input:
        graph.append(list(line.strip()))

    rows = len(graph)
    cols = len(graph[0])

    scenic_score = 3

    for i in range(1, rows - 1):
        for j in range(1, cols - 1):
            scenic_score = max(visible(i, j, graph, rows, cols), scenic_score)

    print(scenic_score)


if __name__ == "__main__":
    main()