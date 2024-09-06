# The JackSnake

The JackSnake is an AI Battlesnake that uses a variety of algorithms to solve a 4-player game of Snake. Each turn runs in under 500 ms.

Its stats can be found here: [JackSnake Stats](https://play.battlesnake.com/leaderboard/standard/m9p909/stats).

Rank at the time of writing (August 2024):
- 36th out of 400 snakes in standard 4-player free-for-all
- 16th out of 300 snakes in duels with 2 players

# High-Level Overview
The Battlesnake uses an algorithm called alpha-beta minimax to search the decision tree of possible moves and determine the best move. In a 4-player game, this setup is called a "Paranoid" algorithm because the alpha-beta pruning assumes that other snakes will always choose the move most likely to kill your snake. The benefit is that you can go much deeper in the decision tree than you could if you didn't use alpha-beta pruning.

A decision tree search will almost never be able to search until the end of the game, so we need a heuristic function. The heuristic function evaluates an unfinished state of the game.  The heuristic function runs hundreds or thousands of times every turn, so it needs to be reasonably performant and provide a good estimation of how the snake is doing.

Minimax Algorithm: [Minimax Algorithm Code](https://github.com/m9p909/jacksnake/blob/main/minimaxplayer/coreplayer/minimaxAlgo.go)

The heuristic function I've found to be the best is called a "Voronoi" function. It uses a breadth-first search to determine how many squares my snake can reach before the other snakes. The more squares my snake controls, the better chance it has of winning. It's combined with a few other metrics like "health." I added a cache to this function because I noticed that many states are repeatedly evaluated.

This algorithm is relatively heavy, but the results are significantly better than any other algorithm. The first time I tested it, this algorithm outperformed my original snake without a game tree search.

Heuristic function: [Heuristic Function Code](https://github.com/m9p909/jacksnake/blob/main/minimaxplayer/evaluator/voronoi_eval.go)

# Performance
Performance is a top concern, and my snake typically executes its moves in about 50-100 ms.

I achieved this performance by profiling the application locally on my laptop to solve issues and using space-efficient data structures so that more data can fit in the CPU caches.

Golang channels handle parallelizing the work. I found that one thread for each possible move works best.
