import unittest

from . import PrimeFactors


class TestPrimeFactors(unittest.TestCase):
    def test_prime_factors(self):
        p = PrimeFactors.PrimeFactors()
        self.assertEqual([], p.generate(1))
        # self.assertEqual([2], p.generate(2))

    def xtest_all_prime_factors(self):
        p = PrimeFactors.PrimeFactors()
        self.assertEqual([], p.generate(1))
        self.assertEqual([2], p.generate(2))
        self.assertEqual([3], p.generate(3))
        self.assertEqual([2, 2], p.generate(4))
        self.assertEqual([5], p.generate(5))
        self.assertEqual([2, 3], p.generate(6))
        self.assertEqual([7], p.generate(7))
        self.assertEqual([2, 2, 2], p.generate(8))
        self.assertEqual([3, 3], p.generate(9))
        self.assertEqual([2, 2, 2, 3, 3, 5], p.generate(360))
        self.assertEqual([2, 3, 5, 7, 13], p.generate(2 * 3 * 5 * 7 * 13))
