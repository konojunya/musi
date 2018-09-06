import * as React from "react"
import styled from "styled-components"

const Header = styled.header`
  width: 100%;
  background-color: black;
  display: flex;
  align-items: center;
  justify-content: center;
  .img {
    width: 30px;
  }
  .logo-text {
    width: 60px;
  }
  img {
    display: block;
    width: 100%;
  }
`

const Head: React.StatelessComponent = () => (
  <Header>
    <div className="img">
      <img src="/images/musi_logo_shape_white.png" alt="musi logo"/>
    </div>
    <div className="logo-text">
      <img src="/images/musi_logo_text_white.png" alt="musi logo"/>
    </div>
  </Header>
)

export default Head;