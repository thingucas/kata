<?php
namespace pdt256\kata\PrimeFactors;

interface PrimeFactorGeneratorInterface
{
    public function generate(int $number): array;
}
