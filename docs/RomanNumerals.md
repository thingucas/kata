
# Roman Numerals

## Rules

Given a positive integer number (eg. 42) determine its Roman numeral representation as a String (eg "XLII").

Roman numerals are expressed by letters of the alphabet:

    I = 1
    V = 5
    X = 10
    L = 50
    C = 100
    D = 500
    M = 1000

1. A letter repeats its value that many times (XXX = 30, CC = 200, etc.). A letter can only be repeated three times.
2. If one or more letters are placed after another letter of greater value, add that amount.
3. If a letter is placed before another letter of greater value, subtract that amount.
    1. Only subtract powers of ten (I, X, or C, but not V or L)
    2. Only subtract one number from another.
    3. Do not subtract a number from one that is more than 10 times greater (that is, you can subtract 1
       from 10 [IX] but not 1 from 20—there is no such number as IXX.)
