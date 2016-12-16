import React from 'react';
import Sound from './Sound.js';
import SoundOverheated from './SoundOverheated.js';

export default class SoundSubList extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            sounds: [],
            mode: "normal"
        };
    }

    componentWillReceiveProps(nextProps) {
        this.setState({
            sounds: nextProps.sounds,
            mode: nextProps.mode
        });
    }

    render() {

        var soundElems = [];

        if (this.state.sounds === undefined) {
            return <h4>Sorry, no data available.</h4>
        }

        this.state.sounds.forEach(item => {
            var soundElem;
            switch (this.state.mode) {
                case "overheated":
                    soundElem = (<SoundOverheated data={item} key={item.SoundFile} />);
                    break;
                default:
                case "normal":
                    soundElem = (<Sound data={item} key={item.SoundFile} />);
                    break;
            }
            soundElems.push(soundElem);
        });

        return (
            <div>
                {soundElems}
            </div>
        );
    }

}
