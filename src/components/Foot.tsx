import * as React from "react"
import * as cookie from "js-cookie";
import styled from "styled-components"

const Footer = styled.footer`
  background-color: black;
  small {
    display: block;
    text-align: center;
    color: white;
    font-size: 12px;
    padding: 5px 0;
  }
`

const Logout = styled.p`
  width: 100%;
  display: block;
  padding: 5px;
  font-size: 12px;
  color: white;
  text-align: center;
`

const Head: React.StatelessComponent = () => (
  <Footer>
    <Logout onClick={() => { cookie.remove("musi-token");location.reload(); }}>ログアウト</Logout>
    <small>&copy; HAL All Right Reserved.</small>
  </Footer>
)

export default Head;