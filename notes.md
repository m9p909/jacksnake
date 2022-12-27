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







