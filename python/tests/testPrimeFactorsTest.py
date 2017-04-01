import unittest

from src import PrimeFactors


class testPrimeFactorsTest(unittest.TestCase):
    def testCreate(self):
        primeFactors = PrimeFactors.PrimeFactors()
