<?php
namespace pdt256\kata\PrimeFactors;

interface PrimeFactorsInterface
{
    public function generate(int $number) : array;
}
