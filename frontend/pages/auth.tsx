import React, {useContext} from "react";
import {ThemeContext} from "../contexts/theme.context";
import {Toggle} from "../components";
import {isLightTheme} from "../helpers/theme.helper";

const Auth = ():JSX.Element => {
    const {theme, setTheme} = useContext(ThemeContext);
    return (
        <Toggle active={isLightTheme(theme)} onClick={() => setTheme && setTheme()}/>
    );
};

export default Auth;