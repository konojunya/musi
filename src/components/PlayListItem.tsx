import * as React from "react"
import styled from "styled-components";

const List = styled.li`
  width: 48vw;
  margin: 5px 0;
  .img {
    width: 100%;
    img {
      display: block;
      width: 100%;
    }
  }
  p {
    text-align: center;
    font-weight: bold;
    font-size: 0.8rem;
    margin-top: 3px;
  }
`

interface Props {
  item;
}

export default class PlayListItem extends React.Component<Props> {
  render() {
    const item = this.props.item;
    return (
      <List onClick={this.handleClick(item.uri)}>
        <div className="img">
          <img src={item.images[0].url} alt=""/>
        </div>
        <p>{item.name}</p>
      </List>
    )
  }

  handleClick = (uri: string) => () => {
    window.open(uri);
  }
}