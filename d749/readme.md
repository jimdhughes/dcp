# Question

This problem was asked by Amazon.

Given a string, find the longest palindromic contiguous substring. If there are more than one with the maximum length, return any one.

For example, the longest palindromic substring of "aabcdcb" is "bcdcb". The longest palindromic substring of "bananas" is "anana".

## Duplicate of d046

# Solving

I broke this out into two problems as this is a mixture of a longest common substring and a palindrome question.

By creating an IsPalindrome function that is independently testable, we can turn this into a longest substring queastion with an extra check for if the string is a palindrome.
I TDD'd this because I'm out of practice with my move to management and I was sure I'd make mistakes. I didn't disappoint :D

I assumed that a single character string is a palindrome in this example. If this was no the case, we could return an empty string if the characters are a single letter instead of just returning the letter.
