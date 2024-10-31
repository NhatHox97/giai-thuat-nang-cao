from collections import deque, defaultdict

def read_input(file_name):
    with open(file_name, 'r') as f:
        data = f.readlines()
    first_line = list(map(int, data[0].split()))
    n = first_line[0]  # số đỉnh
    m = first_line[1]  # số cạnh
    edges = [tuple(map(int, line.split())) for line in data[1:m + 1]]
    return n, edges

def bfs(pair_U, pair_V, dist, adj):
    queue = deque()
    for u in range(1, len(pair_U)):
        if pair_U[u] == 0:  # chưa được ghép
            dist[u] = 0
            queue.append(u)
        else:
            dist[u] = float('inf')
    dist[0] = float('inf')  # đặc biệt cho đỉnh 0

    while queue:
        u = queue.popleft()

        for v in adj[u]:
            if dist[pair_V[v]] == float('inf'):
                dist[pair_V[v]] = dist[u] + 1
                queue.append(pair_V[v])

    return dist[0] != float('inf')

def dfs(pair_U, pair_V, dist, u, adj):
    if u != 0:
        for v in adj[u]:
            if dist[pair_V[v]] == dist[u] + 1:
                if dfs(pair_U, pair_V, dist, pair_V[v], adj):
                    pair_V[v] = u
                    pair_U[u] = v
                    return True
        dist[u] = float('inf')
        return False
    return True

def hopcroft_karp(n, edges):
    adj = defaultdict(list)

    # Xây dựng danh sách kề
    for u, v in edges:
        adj[u].append(v)

    pair_U = [0] * (n + 1)  # ghép cho U
    pair_V = [0] * (n + 1)  # ghép cho V
    dist = [0] * (n + 1)    # khoảng cách

    matching = 0
    while bfs(pair_U, pair_V, dist, adj):
        for u in range(1, len(pair_U)):
            if pair_U[u] == 0:  # chưa được ghép
                if dfs(pair_U, pair_V, dist, u, adj):
                    matching += 1

    return matching

# Đọc dữ liệu từ file và chạy thuật toán
n, edges = read_input('matching_a.txt')
max_matching = hopcroft_karp(n, edges)
print("Số cặp ghép cực đại là:", max_matching)
