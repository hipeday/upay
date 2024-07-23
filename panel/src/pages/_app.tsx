/**
 * APP入口
 *
 * @author jixiangup
 * @since 1.0.0
 */
import type {AppProps} from "next/app";
import {ThemeProvider, type DefaultTheme} from "styled-components";
import GlobalStyle from "@/components/globalstyles";
import {i18n} from "../../i18n-config";
import {useRouter} from "next/router";

export async function generateStaticParams() {
  return i18n.locales.map((locale) => ({ lang: locale }))
}

const defaultLocale = i18n.defaultLocale;

const theme: DefaultTheme = {
  colors: {
    primary: '#000000',
    secondary: '#0070f3',
  }
}


export default function App({ Component, pageProps }: AppProps) {
  const router = useRouter();
  const { lang } = router.query;

  return (
    <>
      <ThemeProvider theme={theme}>
        <GlobalStyle />
        <Component {...pageProps} lang={lang || defaultLocale} />
      </ThemeProvider>
    </>
  )
}
