#ifndef TSP_H
#define TSP_H

#include <vector>
#include "Graph.h"

class TSP {
public:
    int solveTSP(Graph &graph, std::vector<int> &path); // Solve TSP and return minimum cost
    static int VISITED_ALL; // Declaration of the static member variable

private:
    std::vector<std::vector<int>> dp; // DP table
    std::vector<std::vector<int>> prev; // To track the path
    int tspDP(int mask, int pos, Graph &graph); // Recursive DP function
};

#endif // TSP_H
