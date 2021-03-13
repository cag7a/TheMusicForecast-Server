import React, { Component } from "react";
import { WiDaySunny, WiStrongWind, WiRainWind, WiDayCloudy} from 'weather-icons-react';
import "./MainPage.css";

class MainPage extends Component {
    render() {
        return (
            <div id = "MainPage" classname = "page-container">
                <h1>The Music Forecast</h1>
                <div className = "icon-container">
                    <WiDaySunny size={200} color = '#000'/>
                    <WiStrongWind size={200} color = '#000'/>
                    <WiRainWind size={200} color = '#000'/>
                    <WiDayCloudy size={200} color = '#000'/>
                </div>

            </div>
        );
    }
}
export default MainPage; 