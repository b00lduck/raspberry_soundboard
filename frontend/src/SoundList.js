import React from 'react';
import Websocket from 'react-websocket';
import './SoundList.css';
import Sound from './Sound.js';
import SoundOverheated from './SoundOverheated.js';

export default class SoundList extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            AvailableSounds: [],
            OverheatedSounds: []
        };
    }

    handleData(data) {
        let newState = {
            AvailableSounds: [],
            OverheatedSounds: []
        };

        JSON.parse(data).Sounds.forEach(function(x) {
           if (x.Overheated) {
               newState.OverheatedSounds.push(x);
           } else {
               newState.AvailableSounds.push(x);
           }
        });
        this.setState(newState);
    }

    render() {
        return (
            <div className="SoundList">
                {
                    this.state.AvailableSounds.map(item => (
                        <Sound data={item} key={item.SoundFile} />
                    ))
                }
                <br className="clearfix" />
                <hr />
                {
                    this.state.OverheatedSounds.map(item => (
                       <SoundOverheated data={item} key={item.SoundFile} />
                    ))
                }
                <Websocket url="ws://localhost:8080/api/websocket" onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }

}
