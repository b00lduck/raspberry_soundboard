import React from 'react';
import './Sound.css';
import Image from './Image.js';
import Datarow from './Datarow.js';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        this.state = props.data;
    }

    componentWillReceiveProps(props) {
        this.setState(props.data);
    }

    play() {
        fetch('http://pi:8080/api/play/' + this.state.SoundFile);
    }

    render() {
        return (
            <div className="Sound" onClick={this.play.bind(this)}>
                <Image data={this.state} />
                <Datarow data={this.state} />
            </div>
        )
    }
}
