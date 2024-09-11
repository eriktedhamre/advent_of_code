def main():
    with open('example.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()

    nrows = 10
    ncols = 100
    offset = 450
    cave = []

    for i in range(10):
        row = []
        for j in range(100):
            row.append('.')
        cave.append(row)

    #print(cave)
    for line in input:
        split = line.strip().split('->')
        points = []
        for point in split:
            indices = point.split(',')
            points.append((int(indices[1].strip(' ')), int(indices[0].strip(' '))))
        print(points)

    


if __name__ == "__main__":
    main()