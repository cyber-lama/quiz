import React, {useEffect} from "react";
import axios from "axios";

export default function Home():JSX.Element {
    useEffect(() => {
        axios.post('/api/registration', {test: "test"}).then(res => {
            console.log(res)
        }).catch((error)=> {
            console.log(error)
        })
    }, []);
    return (
        <p>test</p>
    );
}
