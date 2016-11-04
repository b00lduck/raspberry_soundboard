import React from 'react';
import './Datarow.css';
import Rainbow from './Rainbow.js';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        var newState = props.data;
        var rainbow = new Rainbow()
        newState.color = '#' + rainbow.colourAt(props.data.Temperature);
        this.state = newState;
    }

    componentWillReceiveProps(props) {
        var newState = props.data;
        var rainbow = new Rainbow()
        newState.color = '#' + rainbow.colourAt(props.data.Temperature);
        this.setState(newState);
    }

    render() {
        return (
            <div className="Datarow">
                <div className="left" style={{background: this.state.color}}>{this.state.Count}x</div>
                <div className="right" style={{background: this.state.color}}>{Math.round(this.state.Temperature * 100) / 100}°</div>
            </div>
        )
    }

}

