import '../styles/globals.scss';
import {AppProps} from "next/dist/shared/lib/router/router";
import React, {useContext} from 'react';
import "@fontsource/roboto";
import styles from "../styles/App.module.scss";
import {WithThemeLayout} from "../layouts/themeLayout/ThemeLayout";
import {ThemeContext} from "../contexts/theme.context";
import Head from "next/head";
import cn from "classnames";
import {css, useTheme} from "@emotion/react";

function MyApp({ Component, pageProps }: AppProps):JSX.Element {
    const themes = useTheme()
    console.log(themes)
  return <>
      <Head>
          <title>Онлайн тестирование</title>
      </Head>
      <div style={{background: themes.colors.background}} className={cn(styles.wrapper)}>
        <Component  {...pageProps} />
      </div>

  </>;
}

export default WithThemeLayout(MyApp);
