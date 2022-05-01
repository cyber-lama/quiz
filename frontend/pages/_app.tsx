import '../styles/globals.scss';
import {AppProps} from "next/dist/shared/lib/router/router";
import React, {useContext} from 'react';
import "@fontsource/roboto";
import cn from "classnames";
import styles from "../styles/Global.module.scss";
import {WithThemeLayout} from "../layouts/themeLayout/ThemeLayout";
import {useTheme} from "@emotion/react";
import {ThemeContext} from "../contexts/theme.context";

function MyApp({ Component, pageProps }: AppProps):JSX.Element {
  const {theme} = useContext(ThemeContext);

  return <>
      <div className={cn(styles.wrapper, {
          [styles.wrapper__light]: theme === 'light',
          [styles.wrapper__dark]: theme === 'dark'
      })}>
        <Component  {...pageProps} />
      </div>

  </>;
}

export default WithThemeLayout(MyApp);
