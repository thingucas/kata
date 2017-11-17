<?php
namespace pdt256\kata\RomanNumerals;

interface RomanNumeralsConverterInterface
{
    public function convertToRoman(int $arabicNumber) : string;
}
