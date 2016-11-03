export default class ColourGradient {

    setGradient(colourStart, colourEnd) {
        this.startColour = this.getHexColour(colourStart);
        this.endColour = this.getHexColour(colourEnd);
    }

    setNumberRange(minNumber, maxNumber) {
        if (maxNumber > minNumber) {
            this.minNum = minNumber;
            this.maxNum = maxNumber;
        }
    }

    colourAt(number) {
        return this.calcHex(number, this.startColour.substring(0,2), this.endColour.substring(0,2))
            + this.calcHex(number, this.startColour.substring(2,4), this.endColour.substring(2,4))
            + this.calcHex(number, this.startColour.substring(4,6), this.endColour.substring(4,6));
    }

    calcHex(number, channelStart_Base16, channelEnd_Base16) {
        var num = number;
        if (num < this.minNum) {
            num = this.minNum;
        }
        if (num > this.maxNum) {
            num = this.maxNum;
        }
        var numRange = this.maxNum - this.minNum;
        var cStart_Base10 = parseInt(channelStart_Base16, 16);
        var cEnd_Base10 = parseInt(channelEnd_Base16, 16);
        var cPerUnit = (cEnd_Base10 - cStart_Base10)/numRange;
        var c_Base10 = Math.round(cPerUnit * (num - this.minNum) + cStart_Base10);
        return this.formatHex(c_Base10.toString(16));
    }

    formatHex(hex) {
        if (hex.length === 1) {
            return '0' + hex;
        } else {
            return hex;
        }
    }

    getHexColour(string) {
        return string.substring(string.length - 6, string.length);
    }

}
