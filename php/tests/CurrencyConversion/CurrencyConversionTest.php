<?php
namespace pdt256\kata\CurrencyConversion;

class CurrencyConversionTest extends \PHPUnit_Framework_TestCase
{
    /** @var CurrencyConversion */
    protected $currencyConversion;

    protected function setUp()
    {
        parent::setUp();
        $this->currencyConversion = new CurrencyConversion();
    }

    // 1) Given 10 USD (U.S. Dollar) convert to 8 EUR (Euro) using 0.8 rate
    public function testConvertUSDtoEUR()
    {
    }

    // 2) Given 30 CHF (Swiss Franc) convert to 3411 JPY (Japanese Yen) using 113.7 exchange rate
    public function testConvertCHFtoJPY()
    {
    }

    // 3) Given 396.5 RUB (Russian Ruble) convert to 4.3615 GBP (British Pound) using 0.011 exchange rate
    public function testConvertRUBtoGBP()
    {
    }
}
