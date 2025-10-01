export class RomanNumerals {
    ROMAN_DIGITS = ["I", "V", "X", "L", "C", "D", "M"]

    public arabicToRoman(arabicNumber: number): string {
        const arabicNumberAsString: string = arabicNumber.toString();
        const arabicNumberLength: number = arabicNumberAsString.length
        let romanNumber = ""

        for (let i = 0; i < arabicNumberLength; i++) {
            const arabicDigit = Number(arabicNumberAsString.charAt(i))
            const potency = arabicNumberLength - (i + 1)
            romanNumber += this.convertToRoman(arabicDigit,  potency)
        }
        return romanNumber;
    }


    public convertToRoman(arabicDigit: number, potency: number) {
        let baseIndex = 2 * potency
        if (arabicDigit === 0) { return "" };
        if(arabicDigit < 4) {
            return this.ROMAN_DIGITS[baseIndex].repeat(arabicDigit)
        }
        if(arabicDigit < 9) {
            const rest = arabicDigit - 5
            if(rest < 0) {
                return this.ROMAN_DIGITS[baseIndex]+ this.ROMAN_DIGITS[baseIndex+1]
            }
            return this.ROMAN_DIGITS[baseIndex+1] + this.ROMAN_DIGITS[baseIndex].repeat(rest)
        }
        if(arabicDigit === 9) {
            return this.ROMAN_DIGITS[baseIndex] + this.ROMAN_DIGITS[baseIndex+2]
        }
    }
}


