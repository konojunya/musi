import * as React from "react"
import styled from "styled-components"

const Wrapper = styled.div`
  width: 100%;
  height: 90vh;
  display: flex;
  justify-content: center;
  align-items: center;
`

const Button = styled.button`
  display: block;
  width: 60%;
  padding: 15px 10px;
  color: white;
  border: 0;
  outline: none;
  border-radius: 10px;
  background-color: black;
`

const ReloadView = () => (
  <Wrapper>
    <Button onClick={() => { location.href = "/login" }}>再読み込みをする</Button>
  </Wrapper>
)
export default ReloadView;