#ifndef GRAPH_H
#define GRAPH_H

#include <vector>

class Graph {
public:
    Graph(int n); // Constructor
    void addEdge(int u, int v, int weight); // Add edge with weight
    int getDistance(int u, int v) const; // Get distance between cities
    int getNumCities() const; // Get the number of cities

private:
    std::vector<std::vector<int>> distance; // Distance matrix
    int numCities; // Number of cities
};

#endif // GRAPH_H
