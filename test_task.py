import sys


def get_filename():
    try:
        return sys.argv[1]
    except IndexError as e:
        raise FileNotFoundError('File name required')


def get_count_line():
    filename = get_filename()

    with open(filename) as f:
        return sum(map(lambda x: 1, f.readlines())), filename


if __name__ == '__main__':
    try:
        print(*get_count_line())
    except Exception as e:
        print(e)
