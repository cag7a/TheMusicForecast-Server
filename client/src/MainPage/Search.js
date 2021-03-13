import React, {useState} from 'react';
import "./MainPage.css";


export default function Search(){
    const [value, setValue] = useState("");

    const handleChange = e => {
        setValue(e.target.value);
    };

    const handleSubmit = e => {
        e.preventDefault();
        alert("You just searched: " + value); 
        setValue('');
    };

    const handleKeypress = e => {
        if(e.keycode == 13){
            handleSubmit();
        }
    };

    return(
        <div className = "input-container">
        <form>
            <input type = "text" 
            className = "search" 
            value = {value}
            onChange = {handleChange}
            onKeyPress={handleKeypress}/> 
            <div className = "btn-container">
                <button className = "btn-search" onClick = {handleSubmit}>
                    Listen To Your Forecast
                </button>
            </div>
        </form>
    </div>
    )
}