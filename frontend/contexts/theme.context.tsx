import {createContext, PropsWithChildren, useState} from 'react';
import {ThemeProvider} from "@emotion/react";

export type ITheme = "light" | "dark";

export interface IThemeContext {
    theme: ITheme
    setTheme?: () => void
}

export const ThemeContext = createContext<IThemeContext>({ theme: "light"});

export const ThemeContextProvider = ({ theme, children }: PropsWithChildren<IThemeContext>): JSX.Element => {
    const [themeState, setThemeState] = useState<ITheme>(theme);
    const setTheme = () => {
        setThemeState(themeState === 'light'? 'dark' : 'light');
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
