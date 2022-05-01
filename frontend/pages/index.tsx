import React, {useContext, useState} from "react";
import {Button, Htag, Ptag, Rating, Tag} from "../components";
import {WithMainLayout} from "../layouts";
import {GetStaticProps} from "next";
import {MenuItem} from "../interfaces/menu.interface";
import menuItems from "../mocks/menu.mock";
import {ThemeContext} from "../contexts/theme.context";

const Home = ({menu}:HomeProps):JSX.Element => {
    // const [menuState, setMenuState] = useState<MenuItem[]>();
    // useEffect(() => {
    //     axios.get('https://jsonplaceholder.typicode.com/users').then(res => {
    //         setMenuState(res.data);
    //     }).catch((error)=> {
    //         console.log(error);
    //     });
    // }, []);
    const [state, setState] = useState(4);
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

            <Rating isEditable rating={state} setRating={setState}/>

            <ul className={'test-sp'}>
                {menu?.map((item, index) => <li key={index}>{item.id.mainCategory}</li>)}
            </ul>
        </>
    );
};

export default WithMainLayout(Home);

export const getStaticProps: GetStaticProps<HomeProps> = async () => {
    // mock
    const menu: MenuItem[] = menuItems;
    return {
        props: {
            menu
        }
    };
};

interface HomeProps extends Record<string, unknown> {
    menu: MenuItem[];
}
