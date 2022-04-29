import '../styles/globals.scss';
import {AppProps} from "next/dist/shared/lib/router/router";
import React from 'react';
import "@fontsource/roboto";
import {ThemeContextProvider} from "../contexts/theme.context";
import cn from "classnames";
import styles from "../styles/Global.module.scss";

function MyApp({ Component, pageProps }: AppProps):JSX.Element {
  return <>
      <ThemeContextProvider theme="light">
          <title>
              Онлайн тестирование
          </title>
          <div className={cn(styles.html__light)}>
              <Component  {...pageProps} />
          </div>
      </ThemeContextProvider>
  </>;
}

export default MyApp;
