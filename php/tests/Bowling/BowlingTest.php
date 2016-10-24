<?php
namespace pdt256\kata\Bowling;

class BowlingTest extends \PHPUnit_Framework_TestCase
{
    public function testSetup()
    {
        $bowling = new Bowling;

        $this->assertTrue($bowling instanceof Bowling);
    }
}
