import ColourGradient from './ColourGradient'

export default class Rainbow {

    constructor() {
        this.minNum = 20;
        this.maxNum = 150;
        this.spectrum = ['00ff00', 'ffff00', 'ff0000', 'ff33ff'];

        let increment = (this.maxNum - this.minNum)/(this.spectrum.length - 1);
        let firstGradient = new ColourGradient();
        firstGradient.setGradient(this.spectrum[0], this.spectrum[1]);
        firstGradient.setNumberRange(this.minNum, this.minNum + increment);
        this.gradients = [ firstGradient ];

        for (var i = 1; i < this.spectrum.length - 1; i++) {
            var colourGradient = new ColourGradient();
            colourGradient.setGradient(this.spectrum[i], this.spectrum[i + 1]);
            colourGradient.setNumberRange(this.minNum + increment * i, this.minNum + increment * (i + 1));
            this.gradients[i] = colourGradient;
        }
    }

    colourAt(number)  {
        var segment = (this.maxNum - this.minNum)/(this.gradients.length);
        var index = Math.min(Math.floor((Math.max(number, this.minNum) - this.minNum)/segment), this.gradients.length - 1);
        return this.gradients[index].colourAt(number);
    }

}

