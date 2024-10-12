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

int main() {
    int n;
    cout << "Enter the number of cities: ";
    cin >> n;

    Graph graph(n);
    cout << "Do you want to manually input the distances or auto-generate them? (m/a): ";
    char choice;
    cin >> choice;

    // Input distance matrix
    if (choice == 'm') {
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
    cout << "The minimum travel cost without return: " << result << endl;
    int returnCost = graph.getDistance(path.back(), 0);
    cout << "The minimum travel cost with return " << totalCost + returnCost << endl;

    cout << "The route to travel is: ";
    // Output the route step by step with costs
    for (size_t i = 0; i < path.size(); ++i) {
        cout << path[i];
        if (i < path.size() - 1) {
            cout << " -> ";
        }
    }
    cout << endl;
    cout << "Details journey: " << endl;
    for (size_t i = 0; i < path.size() - 1; ++i) {
        int from = path[i];
        int to = path[i + 1];
        int cost = graph.getDistance(from, to);
        cout << from << " -> " << to << " (Cost: " << cost << ")\n";
    }
    cout << "Return to starting city: " << 0 << " (Cost: " << returnCost << ")\n";
    return 0;
}
