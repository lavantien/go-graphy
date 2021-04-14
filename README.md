# Graph Theory Briefings

## Graph Types
1. Undirected Graph
2. Directed Graph
3. Weighted Graph
4. Special Graphs *(Tree: Rooted Tree: Arborescence Out-Tree and Anti-Arborescence In-Tree)*
5. Directed Acyclic Graph *(Directed Graph with no cycles. All out-trees are DAGs but not the reverse)*
6. Bipartite Graph *(Dual splittable graph, two colorable or there is no odd length cycle)*
7. Complete Graph *(Kn with n verteces interconnected)*

## Representing Graphs
1. Adjacency Matrix
   | Pros                                          | Cons                                     |
   | --------------------------------------------- | ---------------------------------------- |
   | Space efficient for representing dense graphs | Requires O(V2) space                     |
   | Edge weight lookup is O(1)                    | Iterating over all edges take O(v2) time |
   | Simplest graph representation                 |                                          |
2. Adjacency List
   Example: *A -> [(B,4), (C,1)]*
   | Pros                                           | Cons                                       |
   | ---------------------------------------------- | ------------------------------------------ |
   | Space efficient for representing sparse graphs | Less space efficient for denser graph      |
   | Iterating over all edges is efficient          | Edge weight lookup is O(E)                 |
   | Simplest graph representation                  | Slightly more complex graph representation |
3. Edge List
   Example: *[(C,A,4), (A,C,1), (B,C,6)]*
   | Pros                                           | Cons                                  |
   | ---------------------------------------------- | ------------------------------------- |
   | Space efficient for representing sparse graphs | Less space efficient for denser graph |
   | Iterating over all edges is efficient          | Edge weight lookup is O(E)            |
   | Very simple structure                          |                                       |

## Common Graph Theory Problems
   Is the graph directed or undirected?
   Are the edges of the graph weighted?
   Is the graph likely to be sparse or dense with edges?
   Which structure should I use in order to represent the graph efficiently?
   ### Shortest Path Problem
   **Statement:** Given a weighted graph, find the shortest path of edges from node A to node B.
   **Algorithms:** BFS (unweighted graph), Dijkstra's, Bellman-Ford, Floyd-Warshall, A*, ...
   ### Connectivity
   **Statement:** Does there exist a path between node A and node B?
   **Typical solution:** Use union find data structure or any search algorithm (e.g DFS)
   ### Negative Cycles
   **Statements:** Does my weighted directed graph have any negative cycles? If so, where?
   **Algorithms:** Bellman-Ford and Floyd-Warshall
   ### Strongly Connected Components
   **Statement:** SSCs can be thought of as self-contained cycles within a directed graph where every vertex in a given cycle can reach every other verteces in the same cycle.
   **Algorithms:** Tarjan's and Kosaraju's algorithm
   ### Traveling Saleman Problem
   **Statement:** Given a list of cities and the distances between each pair of cities, what is the shortest possible rout that visits each city exactly once and returns to the original city?
   **Algorithms:** Held-Karp, branch and bound, approximation algorithms (Ant Colony Optimization).
   ### Bridges and Articulation Points
   **Statement:** A bridge / cut edge is any edge in a graph whose removal increases the number of connected components.
   Bridges are important in graph theory because they often hint at weak points, bottlenects or vulnerabilities in a graph.
   ### Minimum Spanning Tree
   **Statement:** A MST is a subset of the edges of a connected weighted graph that connects all the vertices together, without any cycles and with the minimum possible total edge weight.
   **Algorithms:** Kruskal's, Prim's & Boruvka's algorithm
   ### Network Flow: Max Flow
   **Statement:** With an infinite input source how much "flow" can we push through the network?
   Suppose the edges are roads with cars, pipes with water or hallways packed with people.
   Flow represents the volume of water allowed to flow through the pipes, the number of cars the road can sustain in traffic and the maximum amount of people that can navigate through the hallways.
   **Algorithms:** Ford-Fulkerson, Edmonds-Karp & Dinic's algorithm

## Go Implementations
   ### Graph Structure
   Representing graphs in the form of Adjacency List. Also handles initializing (from a Edge List) and outputing.
   ### Depth First Search
