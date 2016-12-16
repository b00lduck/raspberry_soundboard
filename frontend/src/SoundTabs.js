import React from 'react';
import Websocket from 'react-websocket';
import { Tab, Tabs } from 'react-bootstrap';
import './SoundTabs.css';
import SoundSubList from './SoundSubList.js';

export default class SoundTabs extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            AvailableSounds: {},
            OverheatedSounds: [],
            Categories: []
        };
    }

    handleData(data) {
        let newState = {
            AvailableSounds: {},
            OverheatedSounds: [],
            Categories: []
        };

        JSON.parse(data).Sounds.forEach(function(x) {
           if (x.Overheated) {
               newState.OverheatedSounds.push(x);
           } else {
               if (newState.AvailableSounds[x.Category] === undefined) {
                   newState.AvailableSounds[x.Category] = [];
                   newState.Categories.push(x.Category);
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

        var tabs = [];

        this.state.Categories.forEach(item => {
            tabs.push(
                <Tab eventKey={item} title={"Available sounds (" + this.state.AvailableSounds[item].length + ")"}>
                    <SoundSubList mode="normal" sounds={this.state.AvailableSounds[item]} />
                </Tab>
            );
        });

        return (
            <div>
                <Tabs defaultActiveKey="default" id="sound-list-tabs">{tabs}</Tabs>
                <Websocket url="ws://localhost:8080/api/websocket" onMessage={this.handleData.bind(this)}/>
            </div>
        );
    }

}
