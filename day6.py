


def solver(string):

    chars = list(string[0:14])

   

    if len(set(chars)) == 14:
        print(14)
        return

    index = 14
    for i in string[14:]:
        chars.pop(0)
        chars.append(i)


        index += 1
        if len(set(chars)) == 14:
            print(index)
            break



def main():

    with open('input_day6.txt', 'r', encoding='utf-8') as infile:
        string = infile.readline()
    
    string1 = 'mjqjpqmgbljsphdztnvjfqwrcgsmlb'
    string2 = 'bvwbjplbgvbhsrlpgdmjqwftvncz'
    string3 = 'nppdvjthqldpwncqszvftbrmjlhg'
    string4 = 'nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg'
    string5 = 'zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw'

    solver(string)
    solver(string1)
    solver(string2)
    solver(string3)
    solver(string4)
    solver(string5)







if __name__ == "__main__":
    main()