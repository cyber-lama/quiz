import React, {useEffect} from "react";
import axios from "axios";
import {Button, Htag, Ptag, Tag} from "../components";

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
            <Ptag size="l">Большой</Ptag>
            <Ptag size="m">Средний</Ptag>
            <Ptag size="s">Маленький</Ptag>

            <Tag size="s">Chost</Tag>
            <Tag size="m" color='red'>red</Tag>
            <Tag size="s" color='grey'>grey</Tag>
            <Tag size="m" color='green'>green</Tag>
            <Tag size="m" color='primary'>primary</Tag>
        </>
    );
}
