import static java.util.Arrays.asList;
import static org.junit.Assert.*;

import org.junit.Before;
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
    public void generate()
    {
        assertEquals(asList(), primeFactors.generate(0));
        assertEquals(asList(), primeFactors.generate(1));
        assertEquals(asList(2), primeFactors.generate(2));
    }
}
