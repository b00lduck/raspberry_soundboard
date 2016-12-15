import React from 'react';
import Websocket from 'react-websocket';
import { Tabs, Tab } from 'react-bootstrap';
import './SoundTabs.css';
import SoundSubList from './SoundSubList.js';

export default class SoundTabs extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            AvailableSounds: {},
            OverheatedSounds: []
        };
    }

    handleData(data) {
        let newState = {
            AvailableSounds: {},
            OverheatedSounds: []
        };

        JSON.parse(data).Sounds.forEach(function(x) {
           if (x.Overheated) {
               newState.OverheatedSounds.push(x);
           } else {
               if (newState.AvailableSounds[x.Category] === undefined) {
                   newState.AvailableSounds[x.Category] = [];
               }
               newState.AvailableSounds[x.Category].push(x);
           }
        });

        newState.OverheatedSounds.sort(function(a, b) {
           return a.Temperature - b.Temperature;
        });

        this.setState(newState);
    }

    render() {
        return (
            <div>
                <Tabs defaultActiveKey={1} id="sound-list-tabs">
                    <Tab eventKey={1} title={"Available sounds (" + this.state.AvailableSounds.default.length + ")"}>
                        <SoundSubList mode="normal" data={this.state.AvailableSounds.default} />
                    </Tab>
                    <Tab eventKey={2} title={"Overheated sounds (" + this.state.OverheatedSounds.length + ")"}>
                        <SoundSubList mode="overheated" data={this.state.OverheatedSounds} />
                    </Tab>
                </Tabs>
                <Websocket url="ws://localhost:8080/api/websocket" onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }

}
