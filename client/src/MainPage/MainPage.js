import React, { Component } from "react";
import { WiDaySunny, WiStrongWind, WiRainWind, WiDayCloudy} from 'weather-icons-react';
import {FaMusic, FaSpotify} from 'react-icons/fa';
import Search from "./Search";
import "./MainPage.css";


class MainPage extends Component{
    constructor(props) {
        super(props);
        this.state = {
            playlistId: ''
        }
    }
    render() {
        if (this.state.playlistId === ''){
            return (
                <div id = "MainPage" classname = "page-container">
                    <h1>The Music Forecast</h1>
                    <div className = "icon-container">
                        <WiDaySunny size={200} color = '#000'/>
                        <FaSpotify className = "music-icon" size={150} color = '#000'/>
                        <WiDayCloudy  size={200} color = '#000'/>
                        <FaMusic className = "music-icon" size={150} color = '#000'/>
                    </div>
                    <Search 
                    callback = {(playlistId) => this.setState({playlistId})}/>
                </div>
            );
            }
        else{
            return (
                <div id="PlayerPage" className = "player-container">
                    <div className = "playlist">
                        <iframe height = "400px" width = "400px"
                        src={"https://open.spotify.com/embed/playlist/"+ this.state.playlistId}
                        frameborder="0" allowtransparency="true"
                        allow="encrypted-media"></iframe>
                     </div>
                     <div className = "weather-icon">
                        <WiDayCloudy className="weather-icon"  size={200} color = '#000'/>
                     </div>

                </div>
            )
        }
    }
}
export default MainPage; 