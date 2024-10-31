import numpy as np


def read_input(file_name):
    with open(file_name, 'r') as f:
        n = int(f.readline().strip())
        cost_matrix = [list(map(int, f.readline().strip().split())) for _ in range(n)]
    return n, np.array(cost_matrix)


def hungarian_algorithm(cost_matrix):
    from scipy.optimize import linear_sum_assignment

    row_ind, col_ind = linear_sum_assignment(cost_matrix)
    min_cost = cost_matrix[row_ind, col_ind].sum()

    return min_cost, row_ind, col_ind


# Đọc dữ liệu từ file và chạy thuật toán
n, cost_matrix = read_input('matching_weight_c.txt')
min_cost, row_ind, col_ind = hungarian_algorithm(cost_matrix)

print("Tổng chi phí tối ưu là:", min_cost)
print("Cặp ghép (U, V):")
for i in range(len(row_ind)):
    print(f"{row_ind[i] + 1} ghép với {col_ind[i] + 1}")  # +1 để chuyển đổi từ chỉ số 0 sang chỉ số 1
