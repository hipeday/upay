/**
 * APP入口
 *
 * @author jixiangup
 * @since 1.0.0
 */
import type {AppProps} from "next/app";
import {ThemeProvider, type DefaultTheme} from "styled-components";
import GlobalStyle from "@/components/globalstyles";
import {appWithTranslation, i18n} from "next-i18next";
import {useEffect} from "react";

const theme: DefaultTheme = {
  colors: {
    primary: '#000000',
    secondary: '#0070f3',
  },
  fontSizes: {
    primary: '16px',
    mobilePrimary: '14px',
    miniMobilePrimary: '12px',
  },
}


function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <ThemeProvider theme={theme}>
        <GlobalStyle />
        <Component {...pageProps} />
      </ThemeProvider>
    </>
  )
}

export default appWithTranslation(App)