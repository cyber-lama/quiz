import {createContext, PropsWithChildren, useEffect, useLayoutEffect, useState} from 'react';
import {ThemeProvider} from "@emotion/react";
import {getLocalStorage, setLocalStorage} from "../helpers/localstorage.helper";

export type ITheme = "light" | "dark";

export interface IThemeContext {
    theme: ITheme
    setTheme?: () => void
}

export const ThemeContext = createContext<IThemeContext>({ theme: "light"});

export const ThemeContextProvider = ({ theme, children }: PropsWithChildren<IThemeContext>): JSX.Element => {
    const [themeState, setThemeState] = useState<ITheme>(theme);
    // если в LocalStorage нет темы по дефолту добавляем туда light
    // иначе меням
    useEffect(() => {
        if(!getLocalStorage('theme')){
            setLocalStorage('theme', 'light');
        }else{
            setThemeState(getLocalStorage('theme'));
        }
    },[]);

    const setTheme = () => {
        const newTheme = themeState === 'light' ? 'dark' : 'light';
        setLocalStorage('theme', newTheme);
        setThemeState(newTheme);
    };
    const themes = {
        light: {
            colors: {
                font: '#3B434E',
                background: '#FFFAFA',
            },
        },
        dark: {
            colors: {
                font: '#FFFFFF',
                background: '#2F4F4F',
            },
        },
    };

    return <ThemeContext.Provider value={{ theme: themeState, setTheme}}>
        <ThemeProvider theme={themes[themeState]}>
            {children}
        </ThemeProvider>
    </ThemeContext.Provider>;
};
