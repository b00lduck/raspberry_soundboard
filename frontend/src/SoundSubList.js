import React from 'react';
import Sound from './Sound.js';
import SoundOverheated from './SoundOverheated.js';

export default class SoundSubList extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data: props.data
        };
    }

    componentWillReceiveProps(nextProps) {
        this.setState({
            data: nextProps.data
        });
    }

    render() {

        var soundElems = [];

        this.state.data.forEach(item => {
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
