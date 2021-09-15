def print_i(i):
    print(i)

def make_functions():
    flist = []

    for i in [1, 2, 3]:

        flist.append(print_i(i))

    return flist

functions = make_functions()
for f in functions:
    f()