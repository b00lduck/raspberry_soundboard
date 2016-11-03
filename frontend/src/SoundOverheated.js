import React from 'react';
import './SoundOverheated.css';
import Datarow from './Datarow.js';
import Image from './Image.js';

export default class Sound extends React.Component {

    constructor(props) {
        super(props);
        this.state = props.data;
    }

    componentWillReceiveProps(props) {
        this.setState(props.data);
    }

    render() {
        return (
            <div className="SoundOverheated">
                <div>
                    <Image data={this.state} />
                    <Datarow data={this.state} />
                </div>
            </div>
        )
    }
}
