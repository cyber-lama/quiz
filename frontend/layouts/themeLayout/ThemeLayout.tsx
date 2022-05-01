import React, {FunctionComponent} from "react";
import {ThemeContextProvider} from "../../contexts/theme.context";


export const WithThemeLayout =<T extends Record<string, unknown>> (Component: FunctionComponent<T>) => {
    return function WithThemeContext(props: T):JSX.Element{
        return (
            <ThemeContextProvider theme={'light'}>
                    <Component {...props}/>
            </ThemeContextProvider>
        );
    };
};