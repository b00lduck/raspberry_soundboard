import React from 'react';
import './SoundOverheated.css';
import Datarow from './Datarow.js';
import Image from './Image.js';
import Rainbow from './Rainbow.js';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        var newState = props.data;
        var rainbow = new Rainbow();
        newState.color = '#' + rainbow.colourAt(props.data.Temperature);
        this.state = newState;
    }

    componentWillReceiveProps(props) {
        var newState = props.data;
        var rainbow = new Rainbow();
        newState.color = '#' + rainbow.colourAt(props.data.Temperature);
        this.setState(newState);
    }

    render() {
        return (
            <div className="SoundOverheated" style={{filter: "drop-shadow(8px 8px 10px " + this.state.color + ")"}}>
                <div>
                    <Image data={this.state} />
                    <Datarow data={this.state} />
                </div>
            </div>
        )
    }
}
