import React, {useEffect} from "react";
import axios from "axios";
import {Button,Htag} from "../components";

export default function Home():JSX.Element {
    useEffect(() => {
        axios.post('/api/registration', {test: "test"}).then(res => {
            console.log(res)
        }).catch((error)=> {
            console.log(error)
        })
    }, []);
    return (
        <>
            <Htag tag={'h1'}>Test</Htag>
            <Button appearance={'primary'} className={"test"} arrow={'down'}>
                primary
            </Button>
            <Button appearance={'ghost'} arrow={"right"}>
                ghost
            </Button>
        </>
    );
}
