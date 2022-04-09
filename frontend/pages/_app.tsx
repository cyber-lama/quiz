import '../styles/globals.scss';
import {AppProps} from "next/dist/shared/lib/router/router";
import React from 'react';
import "@fontsource/roboto";

function MyApp({ Component, pageProps }: AppProps):JSX.Element {
  return <>
      <title>
          Онлайн тестирование
      </title>
      <Component {...pageProps} />
  </>;
}

export default MyApp;
