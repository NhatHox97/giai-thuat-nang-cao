#include "Graph.h"
#include <iostream>

Graph::Graph(int n) : numCities(n) {
    distance.resize(n, std::vector<int>(n, 0)); // Initialize distance matrix
}

void Graph::addEdge(int u, int v, int weight) {
    distance[u][v] = weight;
    distance[v][u] = weight; // Assuming undirected graph
}

int Graph::getDistance(int u, int v) const {
    return distance[u][v];
}

int Graph::getNumCities() const {
    return numCities;
}
