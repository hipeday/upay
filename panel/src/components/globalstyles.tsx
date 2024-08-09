/**
 * 全局样式
 * @author jixiangup
 * @since 1.0.0
 */

import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
  html,
  body {
    color: ${({ theme }) => theme.colors.primary};
    font-size: ${({ theme }) => theme.fontSizes.primary}
    padding: 0;
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen,
      Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
    
    // mobile style 
    @media (max-width: 768px) {
      font-size: ${({ theme }) => theme.fontSizes.mobilePrimary},
    }

    @media (max-width: 480px) {
      font-size: ${({ theme }) => theme.fontSizes.miniMobilePrimary},
    }
  }

  a {
    color: inherit;
    text-decoration: none;
  }

  * {
    box-sizing: border-box;
  }
`;

export default GlobalStyle;