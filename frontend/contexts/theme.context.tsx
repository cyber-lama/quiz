import { createContext, PropsWithChildren, useState } from 'react';

export type ITheme = "light" | "dark";

export interface IThemeContext {
    theme: ITheme
    setTheme?: (newTheme: ITheme) => void
}

export const ThemeContext = createContext<IThemeContext>({ theme: "light"});

export const ThemeContextProvider = ({ theme, children }: PropsWithChildren<IThemeContext>): JSX.Element => {
    const [themeState, setThemeState] = useState<ITheme>(theme);
    const setTheme = (newTheme: ITheme) => {
        setThemeState(newTheme);
    };

    return <ThemeContext.Provider value={{ theme: themeState, setTheme}}>
        {children}
    </ThemeContext.Provider>;
};