
class File():

    def __init__(self, name, size) -> None:
        self.name = name
        self.size = size



class Dir():
    
    def __init__(self, name, parent) -> None:
        self.parent = parent
        self.name = name
        self.dirs = []
        self.files = []
        self.total_size = 0


dir_sum = 0
def dir_size1(dir):

    global dir_sum
    
    for i in dir.files:
        dir.total_size += i.size
    for i in dir.dirs:
        size = dir_size1(i)
        if size <= 100000:
            dir_sum += size
        dir.total_size += size
    return dir.total_size

def dir_size2(dir):
    
    for i in dir.files:
        dir.total_size += i.size
    for i in dir.dirs:
        size = dir_size2(i)
        dir.total_size += size
    return dir.total_size

    
min_match = 70000000
def find_matching_dirs(dir, goal):

    global min_match
    if dir.total_size > goal:
        min_match = min(dir.total_size, min_match)

    for i in dir.dirs:
        find_matching_dirs(i, goal)
    

def main():
    
    with open('input_day7.txt', 'r', encoding='utf-8') as infile:
        input = infile.readlines()

    root = Dir('/', None)
    cwd = root

    for command in input:
        split = command.split()

        if split[0] == '$':
            if split[1] == 'cd':
                name = split[2]
                if name == '..':
                    cwd = cwd.parent
                else:
                    for i in cwd.dirs:
                        if i.name == name:
                            cwd = i
            else:
                continue
        
        elif split[0] == 'dir':
            name = split[1]
            cwd.dirs.append(Dir(name, cwd))
        else:
            size = int(split[0])
            name = split[1]
            cwd.files.append(File(name, size))
    
    total_size_used = dir_size1(root)
    available = 70000000 - total_size_used
    goal = 30000000 - available
    find_matching_dirs(root, goal)
    print(min_match)
    




if __name__ == "__main__":
    main()