<?php
namespace pdt256\kata\RomanNumerals;

interface RomanNumeralsInterface
{
    public function convertToRoman(int $arabicNumber) : string;
}
