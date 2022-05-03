import '../styles/globals.scss';
import {AppProps} from "next/dist/shared/lib/router/router";
import React, {useContext} from 'react';
import "@fontsource/roboto";
import styles from "../styles/App.module.scss";
import {WithThemeLayout} from "../layouts/themeLayout/ThemeLayout";
import {ThemeContext} from "../contexts/theme.context";
import Head from "next/head";
import cn from "classnames";

function MyApp({ Component, pageProps }: AppProps):JSX.Element {
  const {theme} = useContext(ThemeContext);

  return <>
      <Head>
          <title>Онлайн тестирование</title>
      </Head>
      <div className={cn(styles.wrapper, {
          [styles.wrapper__light]: theme === 'light',
          [styles.wrapper__dark]: theme === 'dark'
      })}>
        <Component  {...pageProps} />
      </div>

  </>;
}

export default WithThemeLayout(MyApp);
