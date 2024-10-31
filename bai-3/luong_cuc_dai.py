from collections import deque

def read_input(file_name):
    with open(file_name, 'r') as f:
        data = f.readlines()
    first_line = list(map(int, data[0].split()))
    n = first_line[0]  # số đỉnh
    m = first_line[1]  # số cạnh
    source = first_line[2]  # đỉnh nguồn
    sink = first_line[3]  # đỉnh đích
    edges = [tuple(map(int, line.split())) for line in data[1:m + 1]]
    return n, m, source, sink, edges

def bfs(capacity, source, sink, parent):
    visited = set()
    queue = deque([source])
    visited.add(source)

    while queue:
        u = queue.popleft()

        for v in range(len(capacity)):
            if v not in visited and capacity[u][v] > 0:  # còn công suất
                queue.append(v)
                visited.add(v)
                parent[v] = u

                if v == sink:
                    return True
    return False

def edmonds_karp(n, source, sink, edges):
    # Khởi tạo ma trận công suất
    capacity = [[0] * (n + 1) for _ in range(n + 1)]
    for u, v, c in edges:
        capacity[u][v] += c  # Nếu có nhiều cạnh thì cộng công suất

    parent = [-1] * (n + 1)
    max_flow = 0

    while bfs(capacity, source, sink, parent):
        path_flow = float('Inf')
        v = sink

        while v != source:
            u = parent[v]
            path_flow = min(path_flow, capacity[u][v])
            v = u

        # Cập nhật công suất của các cạnh và ngược lại
        v = sink
        while v != source:
            u = parent[v]
            capacity[u][v] -= path_flow
            capacity[v][u] += path_flow
            v = u

        max_flow += path_flow

    return max_flow

# Đọc dữ liệu từ file và chạy thuật toán
n, m, source, sink, edges = read_input('maximum_flow_c.txt')
max_flow = edmonds_karp(n, source, sink, edges)
print("Luồng cực đại là:", max_flow)
