import React, {useState} from "react";
import {Button, Htag, Ptag, Rating, Tag} from "../components";
import {WithMainLayout} from "../layouts";
import {GetStaticProps} from "next";
import axios from "axios";
import {UserItem} from "../interfaces/user.interface";

const Home = ({users}:HomeProps):JSX.Element => {
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
                {users?.map(item => <li key={item.id}>{item.name}</li>)}
            </ul>
        </>
    );
};

export default WithMainLayout(Home);

export const getStaticProps: GetStaticProps<HomeProps> = async () => {
    const { data: users } = await axios.get<UserItem[]>("https://jsonplaceholder.typicode.com/users");
    return {
        props: {
            users
        }
    };
};

interface HomeProps extends Record<string, unknown> {
    users: UserItem[];
}
