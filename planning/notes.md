# Making the snake move towards the dots

## Ideas:
### Algo 1:
1. get distance to each dot
2. score of direction = `(sum for each dot (w + h - dist))/(w+h)*(numdots)`

### Algo 2:
1. set score as (w+h-disttonearestdot)

Wait what about the path to the dot. Maybe a graph algorithm? Dijkstras to find shortest path to nearest dot?
Djikstras to find shortest path to each dot, then use the sum like before?

### Algo 3:
1. djikstras shortest path to nearest dot

### Algo 4:
1. get shortest path to each dot and distance of path
2. score of direction = `sum for each dot (w* h - shortestpathdist)/(w*h)*(numdots)`



# Other Algorithms
## Minimax
 Make the decision that minimizes the worst outcome depending on the other sneks. 
 So for each decision estimate that you make and the other snakes make, choose the option that maximizes the minimum score. 

## Alpha Beta Pruning
 used to improve performance of minimax. eliminates an entire branch in the decision tree when score is too low. 
 Still not entirely clear on how this works. Basically eliminates branches that are associated to a really bad decision and dont even consider it. 

## Voronai algorithm
Partition the board into areas where your snake can get first and where other snakes can get first. 
Generate a heuristic for the amount of space you have vs what the other snakes have. 
## BFS
rather than use djikstras, people seem to just use bfs to find the shortest paths

# new ideas based on research

Seems like all these algorithms use a minmax algorithm,
with a heuristic algorithm to determine the value of
each position. 

So if we define a heuristic that works ok for evaluting the 4
current decisions, we could use the same algorithm to 
evaluate the minmax value and the value of other players. 

So we want a function evaluateBoard(state, board,snekid) that receives a board, 
determines the score of the board for a given snek.
The evaluateBoard function should use available heuristics to analyze the board.

Then we can generate a board with each snake move and use heuristics to pick the one with the highest score. 

This algorithm (evaluateBoard) can be expanded with a minmax function to simulate many different boards. 

I think the algorithm should model state as several matrices. A matrix with snake positions, and food. and a matrix with hazards. 

To start with, snake positions and food can be modelled.


# MiniMax with the snakes

I'm not sure how to minimax the game with the snakes because there are multiple opponents. Maybe if I play all the opponents
https://www.aaai.org/Papers/AAAI/1986/AAAI86-025.pdf
https://faculty.cc.gatech.edu/~thad/6601-gradAI-fall2015/Korf_Multi-player-Alpha-beta-Pruning.pdf

It seems to work by alternating the player and having a vector of options. So first the main player goes,  then the opponents go, then the tree can be trimmed. The problem with battlesnake is
that the turns are happening simultaneously. We could check every tree. 

https://coreyja.com/BattlesnakeMinimax/Minimax%20in%20Battlesnake/

There are 2 ways to do a minimax type algorithm: Paranoid and MaxN. In the Paranoid algorithm, we do minmax like it's 2 players and assume all the other snakes are out to get us. 
So the algorithm will analyze the tree as if all the other snakes are trying to minimize our score. 

In the MaxN algorithm we assume that all the other snakes are trying to maximize their own score. So we pick the best choice assuming all other snakes are acting rationally. 

To use alpha beta, we need pruning. MaxN should be more accurate, but much slower. Alpha beta pruning does not work in maxN. 

For this algorithm, I think I'll use alpha beta because It will have better performance.


