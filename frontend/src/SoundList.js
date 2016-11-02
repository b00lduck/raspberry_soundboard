import React from 'react';
import Sound from './Sound.js';
import Websocket from 'react-websocket';

export default class SoundList extends React.Component {

    handleData(data) {
        let result = JSON.parse(data);
        console.log(result);
        //this.setState({count: this.state.numPlayed + result.movement});
    }

    render() {
        return (
            <div className="SoundList">
                <Sound name="sound1" />
                <Sound name="sound2" />
                <Websocket url='ws://localhost:8080/api/websocket' onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }
}
