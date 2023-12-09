# Advent of Code 2023

My solutions to the puzzles.
All programs read the puzzle from standard input and print the solution to standard output, just like good old [Unix filters](https://en.wikipedia.org/wiki/Filter_(software)) are supposed to do.

## Day 1

**Problem.**
You are given lines of text that contain digits and letters.
For each line, you are asked to find the first and the last digit and combine them into a number.
Finally, you sum all the numbers.

```
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

In part 2, you find out that digits can be found as words as well.

```
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
```

**Comment.**
Initially, part 2 worked on tests but failed on input, and I couldn't understand why.
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

```
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
```

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

```
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
```

**Comment.** A nice little recursive function solved the problem.
For once, I also manage to write proper tests.
Part 2 was a simple flip in the logic that required just a couple of minutes for implementation.
Overall, today's solution came smoothly and I feel satisfied.
