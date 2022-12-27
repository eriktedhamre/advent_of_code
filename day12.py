

alphabet = {'a': 1,
            'b': 2,
            'c': 3,
            'd': 4,
            'e': 5,
            'f': 6,
            'g': 7,
            'h': 8,
            'i': 9,
            'j': 10,
            'k': 11,
            'l': 12,
            'm': 13,
            'n': 14,
            'o': 15,
            'p': 16,
            'q': 17,
            'r': 18,
            's': 19,
            't': 20,
            'u': 21,
            'v': 22,
            'w': 23,
            'x': 24,
            'y': 25,
            'z': 26
            }

def main():
    with open('example.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()
    
    graph = []
    for line in input:
        graph.append(list(line.strip()))

    nrows = len(graph)
    ncol = len(graph[0])
    src 
    
    print(graph)

if __name__ == "__main__":
    main()