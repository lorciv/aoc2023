# Advent of Code 2023

My solutions to the puzzles.

## Day 1

Initially, part 2 worked on tests but failed on input. I couldn't understand why.
Then on r/adventofcode I read that the program needs to take into account overlapping words.
It was a simple change in the code, but left me disheartened.

## Day 2

I found input parsing to be a bit too elaborated for being just the second day of the challenge.
Postponed part 2 for now. Omitted tests as well.

## Day 3

Super tedious!
Part 2 skipped for now.

## Day 4

**Problem.**
You are given a list of scratchcards.
In part 1, you are asked to compute the score for each card and sum them all together.
In part 2, cards make you win copies of other cards.
Instead of computing the score of each card, you count the number of cards you end up with.

**Comment.**
Finally an easier task!
I just parsed the cards, put them into a list, and then do all the calculations, which are straightforward.
I represent multiple copies of the same card with an additional field "count" inside the card struct (instead of creating multple card objects).

## Day 5

Not very happy with the solution, but it works.
I read on r/advendofcode that part 2 is hard. No problem, I'll skip it.

## Day 6

Doable, at least in part 1.

## Day 9

**Problem.** Given a sequence of numbers, predict the next one.
The algorithm to predict the next number is given by the assignment, and it is recursive.
Part 2 asks to predict the value before the first one, following a similar logic.

**Comment.** A nice little recursive function solved the problem.
For once, I also manage to write proper tests.
Part 2 was a simple flip in the logic that required just a couple of minutes for implementation.
Overall, today's solution came smoothly and I feel satisfied.
