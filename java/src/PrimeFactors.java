import java.util.ArrayList;
import java.util.List;

class PrimeFactors
{
    List<Integer> generate(Integer number)
    {
        List<Integer> factors = new ArrayList<>();

        if (number == 2) {
            factors.add(2);
        }

        return factors;
    }
}
