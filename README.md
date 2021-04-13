### Graph Theory Briefings

## Graph Types
1. Undirected Graph
2. Directed Graph
3. Weighted Graph
4. Special Graphs (Tree: Rooted Tree: Arborescence Out-Tree and Anti-Arborescence In-Tree)
5. Directed Acyclic Graph (Directed Graph with no cycles. All out-trees are DAGs but not the reverse)
6. Bipartite Graph (Dual splittable graph, two colorable or there is no odd length cycle)
7. Complete Graph (Kn with n verteces interconnected)

## Representing Graphs
1. Adjacency Matrix
   | Pros                                          | Cons                                     |
   | --------------------------------------------- | ---------------------------------------- |
   | Space efficient for representing dense graphs | Requires O(V2) space                     |
   | Edge weight lookup is O(1)                    | Iterating over all edges take O(v2) time |
   | Simplest graph representation                 |                                          |
2. Adjacency List
   Example: A -> [(B,4), (C,1)]
   | Pros                                           | Cons                                       |
   | ---------------------------------------------- | ------------------------------------------ |
   | Space efficient for representing sparse graphs | Less space efficient for denser graph      |
   | Iterating over all edges is efficient          | Edge weight lookup is O(E)                 |
   | Simplest graph representation                  | Slightly more complex graph representation |
3. Edge List
   Example: [(C,A,4), (A,C,1), (B, C, 6)]
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
   # Shortest Path Problem
   Statement: Given a weighted graph, find the shortest path of edges from node A to node B.
   Algorithms: BFS (unweighted graph), Dijkstra's, Bellman-Ford, Floyd-Warshall, A*, ...
   