<?php
namespace pdt256\kata\RomanNumerals;

interface RomanNumeralsConverterInterface
{
    public function arabicToRoman(int $arabicNumber) : string;
}
