import React, {useContext} from "react";
import {ThemeContext} from "../contexts/theme.context";

const Auth = ():JSX.Element => {
    const {setTheme} = useContext(ThemeContext);
    return <p onClick={() => setTheme && setTheme()}>test</p>;
};

export default Auth;