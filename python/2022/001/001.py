def do_magic(file):
    with open(file) as f:
        cur = 0
        vals = []

        lines = f.readlines()
        for line in lines:
            if line == '\n':
                vals.append(cur)
                cur = 0
                continue

            cur += int(line)

        if cur != 0:
            vals.append(cur)

        vals.sort()

        print(vals[len(vals)-1])
        print(vals[len(vals)-1]+vals[len(vals)-2]+vals[len(vals)-3])

do_magic('./example.txt')
do_magic('./input.txt')