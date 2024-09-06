# The JackSnake

I started the jacksnake over christmas break when I discovered Battlesnake! It’s a game where you write an AI snake and compete against other AI snakes.
# Basics

The first step was to make the snake move towards the food. I used a simple breadth first search to find the shortest path to the food. This worked well, but the snake still killed itself because it would use too much space. The snake didn’t know to stay along the border or fold into itself. It also chased food, even when it had full or nearly full health!.
# Research

Before doing anything else, I read up on the most common algorithms used in AI snakes. This is what I discovered:
## Minimax

a lot of snakes were tagged as minimax or alpha beta snakes. Minimax is an adversarial search algorithm that is used to find the best move for a player. It works by generating the game tree, and alternating between evaluating the best move for the player and the enemy.

Minimax is a recursive function. The base case is when a certain depth is reached or the game is over. Pre-recursion it generates all the possible moves for the current player. Then it calls itself on each of those moves. If the row is a maximizing row (current player), then it return the maximum value for the current player. If it’s a minimizing row (enemy), then it returns the minimum child value. It simulates a game where the opponent is trying to minimize the score and the current player is trying to maximize the score. And chooses the move that maximizes the score.
## Alpha-Beta Pruning

This algorithm is an extension or technique for minimax. If the score isn’t strong enough to change the decision at a higher node, then it doesn’t need to be evaluated. For example if a min node chose a score of 8, then and one the remaining max nodes receives a value higher than 8, then we don’t need to keep evaluating other nodes because we know that the previous min node won’t chose it.
## Voronai

Partition the board into areas you’re snake can get to first, and the areas the enemy snake can get to first. We can use this information as a heuristic in our minimax algorithm. Using this heuristic makes our snake more aggressive.
## Implementation

I started building the snake in Go, because I like the language. It’s around as fast as java, so it should be plenty fast unless I get really competitive and it compiles to a single executable. No scripting runtime, or JVM. Just a single binary.

Thus far, I’ve implemented a basic heuristic for food and space. I need to add other heursitics so my snake can get smarter.
Minimax Implementation

I’m working on the minimax algorithm. The problem with the minimax algorithm is that it’s made for 2 player games.

To make it work with a 4 player games there are 2 approachs: Paranoid, and MaxN
## Paranoid

paranoid is when you assume that the enemy is trying to minimize your own score. So every opponent is a min node in the minmax tree. This method provides more depth because we can use the alpha beta algorithm.
## MaxN

Max^N is where we assume all other snakes are trying to maximize their own score. This method makes more sense and is probably more accurate for multiplayer games but alpha-beta pruning doesn’t work. So we need to try every possible game state
untitled page
