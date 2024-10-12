#include <iostream>
#include <vector>

using namespace std;

/**
* Gôm tất cả file nhỏ lại thành main-c-free để dễ dàng chạy chương trình trên C-Free
*/

class Graph {
public:
    Graph(int n); // Constructor
    void addEdge(int u, int v, int weight); // Add edge with weight
    int getDistance(int u, int v) const; // Get distance between cities
    int getNumCities() const; // Get the number of cities

private:
    std::vector<std::vector<int> > distance; // Distance matrix
    int numCities; // Number of cities
};

class TSP {
public:
    int solveTSP(Graph &graph, std::vector<int> &path); // Solve TSP and return minimum cost
    static int VISITED_ALL; // Declaration of the static member variable

private:
    std::vector<std::vector<int> > dp; // DP table
    std::vector<std::vector<int> > prev; // To track the path
    int tspDP(int mask, int pos, Graph &graph); // Recursive DP function
};

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
    dp = std::vector<std::vector<int> >(1 << n, std::vector<int>(n, -1));
    prev = std::vector<std::vector<int> >(1 << n, std::vector<int>(n, -1)); // Initialize previous cities tracker

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



void displayGraph(const vector<vector<int> >& graph) {
    cout << "City Distance Map:\n";
    for (size_t i = 0; i < graph.size(); ++i) {
        for (size_t j = 0; j < graph[i].size(); ++j) {
            cout << graph[i][j] << "\t";
        }
        cout << endl;
    }
}

/**
 * X? lý bài toán ngu?i bán hàng b?ng phu?ng pháp quy ho?ch d?ng
 * K?t h?p v?i bitmask d? dánh d?u các thành ph? dã qua 1 cách hi?u qu? hon
 * Tham kh?o t?i ngu?n: https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
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
    vector<vector<int> > distanceMatrix(n, vector<int>(n));
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
