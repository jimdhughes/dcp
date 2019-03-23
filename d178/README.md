# Question
Good morning! Here's your coding interview problem for today.

This problem was asked by Two Sigma.

Alice wants to join her school's Probability Student Club. Membership dues are computed via one of two simple probabilistic games.

The first game: roll a die repeatedly. Stop rolling once you get a five followed by a six. Your number of rolls is the amount you pay, in dollars.

The second game: same, except that the stopping condition is a five followed by a five.

Which of the two games should Alice elect to play? Does it even matter? Write a program to simulate the two games and calculate their expected value.

# Answer
Your chances of landing a 5 on any given roll is 1/6. Your chances of landing a 6 on any given roll is 1/6. Your chances of landing a 5 on a roll and a 6 on another roll is 1/6 * 1/6 so 1/36.  I'd expect that this would cost about $36 on average.

Your chances of rolling a 5 directly after rolling a 5 should be the same.  You have a 1/6 chance of rolling a 5 on any given roll and a 1/6 chance of rolling another 5 on any given roll threfore I'd expect that the logic holds true.

# Simulation
With 1,000,000 simulations, my answer for the first scenario holds true for rolling a 5 and then a 6.

This was not true for the second scenario however.  With 1,000,000 iterations the average cost is 41.

# Learning
Forgetting that permutations and combinations matter, my original logic does not hold true.  For combinations, the 1/6 * 1/6 logic holds true.  However when we are look