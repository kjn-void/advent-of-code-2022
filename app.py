dir = {"U": complex(0,1), "D": complex(0,-1), "L": complex(-1,0), "R": complex(1,0)}
parts = [complex(0,0)] * 10
visited = [set(), set(), set(), set(), set(), set(), set(), set(), set(), set(), set()]

for move in open("inputs/day9.txt").readlines():
    m = move.strip().split()
    for _ in range(int(m[1])):
        parts[0] += dir[m[0]]
        for i in range(1, len(parts)):
            if abs(diff := parts[i - 1] - parts[i]) > 1.5:
                parts[i] += complex(diff.real / abs(diff.real) if diff.real else 0,
                                    diff.imag / abs(diff.imag) if diff.imag else 0) 
            visited[i].add(parts[i])
print([len(visited[i]) for i in [1,9]])
