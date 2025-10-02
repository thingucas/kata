export class RomanNumerals {
    ROMAN_DIGITS = ["I", "V", "X", "L", "C", "D", "M"]

    public arabicToRoman(arabicNumber: number): string {
        const arabicNumberAsString: string = arabicNumber.toString();
        let romanNumber = ""

        for (let i = 0; i < arabicNumberAsString.length; i++) {
            const arabicDigit = Number(arabicNumberAsString.charAt(i))
            const potency = arabicNumberAsString.length - (i + 1)
            romanNumber += this.convertDecimalPotency(arabicDigit,  potency)
        }
        return romanNumber;
    }


    private convertDecimalPotency(arabicDigit: number, potency: number) {
        let baseIndex = 2 * potency
        if(this.isNumberWithSubstraction(arabicDigit)) {
            return this.getSubtrahend(baseIndex) + this.getMinuend(arabicDigit,baseIndex)
        }
        return this.getFirstSummand(arabicDigit,baseIndex) + this.getSecondSummand(arabicDigit,baseIndex)
    }

    private isNumberWithSubstraction(arabicDigit: number): boolean {
        return arabicDigit % 5 === 4
    }

    private getSubtrahend(baseIndex: number): string {
        return this.ROMAN_DIGITS[baseIndex]
    }

    private getMinuend(arabicDigit: number, baseIndex: number): string {
        return this.ROMAN_DIGITS[baseIndex+1+Math.floor(arabicDigit/5)];
    }

    private getSecondSummand(arabicDigit: number, baseIndex: number): string {
        return this.ROMAN_DIGITS[baseIndex].repeat(arabicDigit % 5)
    }

    private getFirstSummand(arabicDigit: number, baseIndex: number): string {
        return arabicDigit >= 5 ? this.ROMAN_DIGITS[baseIndex+1] : ""
    }
}


