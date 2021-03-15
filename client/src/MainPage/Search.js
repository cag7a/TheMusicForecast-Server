import React, {useState} from 'react';
import "./MainPage.css";


export default function Search(props){
    const [value, setValue] = useState("");

    const handleChange = e => {
        setValue(e.target.value);
    };

    const handleSubmit = e => {
        e.preventDefault();
        fetch(`/api/playlist?search=${value}`).then((data) =>
         data.json()).then((data) => props.callback(data.playlist));
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