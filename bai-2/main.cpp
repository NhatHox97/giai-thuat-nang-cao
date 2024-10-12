#include <iostream>
#include <vector>
#include "Graph.h"
#include "TSP.h"

using namespace std;

void displayGraph(const vector<vector<int>>& graph) {
    cout << "City Distance Map:\n";
    for (size_t i = 0; i < graph.size(); ++i) {
        for (size_t j = 0; j < graph[i].size(); ++j) {
            cout << graph[i][j] << "\t";
        }
        cout << endl;
    }
}

/**
 * Xữ lý bài toán người bán hàng bằng phướng pháp quy hoạch động
 * Kết hợp với bitmask để đánh dấu các thành phố đã qua 1 cách hiệu quả hơn
 * Tham khảo tại nguồn: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
 */
int main() {
    int n;
    cout << "Enter the number of cities: ";
    cin >> n;

    Graph graph(n);
    cout << "Do you want to auto-generate city? (y/n): ";
    char choice;
    cin >> choice;

    // Input distance matrix
    if (choice == 'n') {
        cout << "Enter the distance matrix (0 for same city):\n";
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                int distance;
                cin >> distance;
                graph.addEdge(i, j, distance);
            }
        }
    } else {
        srand(time(0)); // Seed for random number generation
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (i == j) {
                    graph.addEdge(i, j, 0); // Distance to itself is 0
                } else {
                    graph.addEdge(i, j, rand() % 100 + 1); // Random distance between 1 and 100
                }
            }
        }
    }

    // Display the graph
    vector<vector<int>> distanceMatrix(n, vector<int>(n));
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < n; ++j) {
            distanceMatrix[i][j] = graph.getDistance(i, j);
        }
    }
    displayGraph(distanceMatrix); // Display the distance matrix

    TSP tspSolver;
    vector<int> path; // To store the route
    int result = tspSolver.solveTSP(graph, path); // Solve TSP
    int totalCost = result; // Start with the minimum cost

    // Display results
    cout << "The minimum travel cost: " << result << endl;

    cout << "The route to travel is: ";
    // Output the route step by step
    for (size_t i = 0; i < path.size(); ++i) {
        cout << path[i];
        if (i < path.size() - 1) {
            cout << " -> ";
        }
        if (i == path.size() -1)
        {
            cout << " -> 0";
        }
    }
    cout << endl;
    cout << "-------------------------------" << endl;
    cout << "Details journey: " << endl;

    int cumulativeCost = 0; // To track the cumulative cost
    cout << "Current Cost: " << cumulativeCost << endl;
    for (size_t i = 0; i < path.size() - 1; ++i) {
        int from = path[i];
        int to = path[i + 1];
        int cost = graph.getDistance(from, to);
        cumulativeCost += cost; // Update cumulative cost
        cout << from << " -> " << to << " (Cost: " << cost << ", Current Cost: " << cumulativeCost << ")\n";
    }

    int returnCost = graph.getDistance(path.back(), 0);
    cumulativeCost += returnCost; // Update cumulative cost for return
    cout << "Return to starting city: " << 0 << " (Cost: " << returnCost << ", Total Travel Cost: " << cumulativeCost << ")\n";

    return 0;
}
