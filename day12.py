import queue

alphabet = {'E':26,
            'S':1,
            'a': 1,
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

neighbors = [(1, 0),(0, 1),(-1, 0),(0, -1)]

def bfs(graph, src, visited):

    nrows = len(graph)
    ncols = len(graph[0])

    q = queue.Queue()

    visited[src[0]][src[1]] = 1

    q.put((0, 0, 0))

    while not q.empty():
        u = q.get()

        char = graph[u[0]][u[1]]
        c_value = alphabet[char]
        if char == 'E':
            return u[2]
        
        for mod in neighbors:
            n_row = u[0] + mod[0]
            n_col = u[1] + mod[1]
            if n_row > -1 and n_row < nrows and n_col > -1 and n_col < ncols and not visited[n_row][n_col]:
                n_value = alphabet[graph[n_row][n_col]]
                
                if c_value + 1 >= n_value:
                    q.put((n_row, n_col, u[2] + max(1,max(c_value, n_value) - min(c_value, n_value))))
                    visited[n_row][n_col] = 1



def main():
    with open('input.txt', 'r', encoding='utf-8') as f:
        input = f.readlines()
    
    graph = []
    for line in input:
        graph.append(list(line.strip()))

    nrows = len(graph)
    ncols = len(graph[0])

    src = (0, 0, 0)
    dst = 'E'
    visited = []
    for i in range(nrows):
        row = []
        for j in range(ncols):
            row.append(0)
        visited.append(row)
    
    print(bfs(graph, src, visited))

if __name__ == "__main__":
    main()