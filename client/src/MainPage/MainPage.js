import React, { Component } from "react";
import { WiDaySunny, WiStrongWind, WiRainWind, WiDayCloudy} from 'weather-icons-react';
import {FaMusic, FaSpotify} from 'react-icons/fa';
import Search from "./Search";
import "./MainPage.css";



class MainPage extends Component {
    render() {
        return (
            <div id = "MainPage" classname = "page-container">
                <h1>The Music Forecast</h1>
                <div className = "icon-container">
                    <WiDaySunny size={200} color = '#000'/>
                    <FaSpotify className = "music-icon" size={150} color = '#000'/>
                    <WiDayCloudy  size={200} color = '#000'/>
                    <FaMusic className = "music-icon" size={150} color = '#000'/>
                </div>
                <Search/>
            </div>
        );
    }
}
export default MainPage; 