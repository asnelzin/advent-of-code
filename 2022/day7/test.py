sample = open("input.txt")


class Node:
    def __init__(self, name, parent, size=0):
        self.name = name
        self.parent = parent
        self.dirs = {}
        self.files = {}
        self.size = size


root = Node("/", None)
pwd = root

print("- / (dir)")

for c in sample.read().split("$")[1:]:
    c = c.strip("\n")
    c = c.strip()
    if c.startswith("ls"):
        c = c.split("\n")
        ll = c[1:]
        for item in ll:
            t, name = item.split(" ")
            if t == "dir":
                pwd.dirs[name] = Node(name, pwd)
            else:
                pwd.files[name] = Node(name, pwd, size=int(t))
    elif c.startswith("cd"):
        command, arg = c.split(" ")
        if arg == "/":
            pwd = root
        elif arg == "..":
            pwd = pwd.parent
        else:
            pwd = pwd.dirs[arg]


total = 0
st = [root]

sol = []


def calc(node):
    st = [node]

    st.extend(node.dirs.values())

    i = 1
    while i < len(st):
        dir = st[i]
        print(dir.name)
        if dir.dirs:
            st.extend(dir.dirs.values())
        i += 1

    while st:
        node = st.pop()

        node.size = sum([f.size for f in node.files.values()]) + sum(
            [d.size for d in node.dirs.values()]
        )

        print(node.size)

        if node.size != 0:
            sol.append(node.size)


calc(root)
print(list(sorted(sol)))

for idx, size in enumerate(list(sorted(sol))):
    if size >= 30000000 - (70000000 - root.size):
        print(list(sorted(sol))[idx])
        break
