import React, {useContext} from "react";
import {ThemeContext} from "../contexts/theme.context";
import {Toggle} from "../components";

const Auth = ():JSX.Element => {
    const {theme, setTheme} = useContext(ThemeContext);
    return (
        <Toggle active={theme === 'light'} onClick={() => setTheme && setTheme()}/>
    );
};

export default Auth;