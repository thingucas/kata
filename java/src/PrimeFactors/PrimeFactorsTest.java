package PrimeFactors;

import static java.util.Arrays.asList;
import static org.junit.Assert.*;

import org.junit.Before;
import org.junit.Ignore;
import org.junit.Test;

public class PrimeFactorsTest
{
    private PrimeFactors primeFactors;

    @Before
    public void setup()
    {
        primeFactors = new PrimeFactors();
    }

    @Test
    public void primeFactors()
    {
        assertEquals(asList(), primeFactors.generate(1));
        //assertEquals(asList(2), primeFactors.generate(2));
    }

    @Test
    @Ignore
    public void allPrimeFactors()
    {
        assertEquals(asList(2), primeFactors.generate(2));
        assertEquals(asList(3), primeFactors.generate(3));
        assertEquals(asList(2, 2), primeFactors.generate(4));
        assertEquals(asList(5), primeFactors.generate(5));
        assertEquals(asList(2, 3), primeFactors.generate(6));
        assertEquals(asList(7), primeFactors.generate(7));
        assertEquals(asList(2, 2, 2), primeFactors.generate(8));
        assertEquals(asList(3, 3), primeFactors.generate(9));
        assertEquals(asList(2, 2, 2, 3, 3, 5), primeFactors.generate(360));
        assertEquals(asList(2, 3, 5, 7, 13), primeFactors.generate(2 * 3 * 5 * 7 * 13));
    }
}
