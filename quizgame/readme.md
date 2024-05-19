# Quiz Game
This quiz game reads in a CSV file with a format of "question, answer" with a timer. You can set which csv files to read by adding a `-csv=` flag and what the timer length for each question should be by adding a `-limit=' flag in the execute statement.

Sample CSV:

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Sample output with a timeout using `./quiz -limit=5`
```
Problem #1: 5+5 =
10
Problem #2: 7+3 =
10
Problem #3: 1+1 =
2
Problem #4: 8+3 =

You scored 3 out of 13.
```
