import {RomanNumerals} from './roman-numerals';

let romanNumerals;

describe('RomanNumeral Test Module', () => {
    beforeEach(() => {
        romanNumerals = new RomanNumerals();
    });

    afterEach(() => {
        romanNumerals = null;
    });

    it('Converts the value to the roman numeral equivalent.', () => {
        expect(romanNumerals.arabicToRoman(468)).toEqual('CDLXVIII')
        expect(romanNumerals.arabicToRoman(2)).toEqual('II')
    });
});
