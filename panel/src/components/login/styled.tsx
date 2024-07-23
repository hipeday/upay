/**
 * 登录页面样式
 *
 * @author jixiangup
 * @since 1.0.0
 */

import styled from "styled-components";

const Main = styled.main`
  display: flex;
  align-items: flex-start;
  flex-direction: column;
`

const HeadContainer = styled.div`
  margin: 0;
  width: 100%;
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
`;

const ContentContainer = styled.div`
  display: flex;
  .content-background {
    width: 827px;
    height: 650px;
  }
`

const FormContainer = styled.div`
  border: 1px solid #878787;
  width: 505px;
  height: 701px;
  border-radius: 10px;
`

export { HeadContainer, FormContainer, ContentContainer, Main };