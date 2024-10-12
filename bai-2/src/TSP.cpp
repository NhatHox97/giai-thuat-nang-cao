#include "TSP.h"
#include <climits>

// Define the static member variable
int TSP::VISITED_ALL = 0;

int TSP::tspDP(int mask, int pos, Graph &graph) {
    int n = graph.getNumCities();

    // Base case: all cities have been visited
    if (mask == VISITED_ALL) {
        return graph.getDistance(pos, 0); // Return to the starting city
    }

    // Check if we already have a computed value
    if (dp[mask][pos] != -1) {
        return dp[mask][pos];
    }

    // Try all possible cities to visit next
    int ans = INT_MAX;
    for (int city = 0; city < n; ++city) {
        // Check if city has been visited
        if ((mask & (1 << city)) == 0) {
            int newAns = graph.getDistance(pos, city) + tspDP(mask | (1 << city), city, graph);
            if (newAns < ans) {
                ans = newAns;
                prev[mask][pos] = city; // Track the path
            }
        }
    }

    // Store the result in the DP table and return it
    return dp[mask][pos] = ans;
}

int TSP::solveTSP(Graph &graph, std::vector<int> &path) {
    int n = graph.getNumCities();

    // Initialize the DP table with -1 (meaning uncomputed)
    dp = std::vector<std::vector<int>>(1 << n, std::vector<int>(n, -1));
    prev = std::vector<std::vector<int>>(1 << n, std::vector<int>(n, -1)); // Initialize previous cities tracker

    // Bitmask representing all cities visited (e.g., for 4 cities it's 1111 in binary)
    VISITED_ALL = (1 << n) - 1;

    // Start the DP from the first city (position 0) with only the first city visited (mask = 1)
    int minCost = tspDP(1, 0, graph);

    // Reconstruct the path
    path.clear();
    int mask = 1; // Start with only the first city visited
    int currentCity = 0;

    path.push_back(currentCity); // Start from city 0
    while (mask != VISITED_ALL) {
        currentCity = prev[mask][currentCity];
        mask |= (1 << currentCity);
        path.push_back(currentCity);
    }

    return minCost;
}
