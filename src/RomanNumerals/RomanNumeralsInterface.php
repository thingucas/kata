<?php
namespace pdt256\kata\RomanNumerals;

interface RomanNumeralsInterface
{
    public function convertArabic(int $arabicNumber) : string;
}
