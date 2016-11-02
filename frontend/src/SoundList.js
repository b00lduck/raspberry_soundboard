import React from 'react';
import Websocket from 'react-websocket';
import './SoundList.css';
import Sound from './Sound.js';

export default class SoundList extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            Sounds: []
        };
    }

    handleData(data) {
        this.setState(JSON.parse(data));
    }

    render() {
        return (
            <div className="SoundList">
                {this.state.Sounds.map(item => (
                       <Sound data={item} key={item.SoundFile} />
                ))}
                <Websocket url="ws://localhost:8080/api/websocket" onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }

}
