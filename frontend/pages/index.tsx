import React, {useEffect, useState} from "react";
import axios from "axios";
import {Button, Htag, Ptag, Rating, Tag} from "../components";

export default function Home():JSX.Element {
    const [state, setState] = useState(4)
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

            <Rating isEditable={true} rating={state} setRating={setState}/>
        </>
    );
}
