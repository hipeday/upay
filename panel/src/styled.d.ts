/**
 * 样式配置
 *
 * @author jixiangup
 * @since 1.0.0
 */

import "styled-components";

declare module "styled-components" {
  export interface DefaultTheme {
    colors: {
      primary: string;
      secondary: string;
    };
  }
}